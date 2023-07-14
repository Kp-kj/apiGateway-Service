package user

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/userclient"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ResponseData struct {
	Code int `json:"code"`
	Data struct {
		AccessToken string `json:"access_token"`
		Name        string `json:"name"`
		UserID      string `json:"user_id"`
		Username    string `json:"username"`
		Avatar      string `json:"profile_image_url"`
	} `json:"data"`
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 创建jwt token 逻辑
func (l *UserLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// 创建用户jwt token
func (l *UserLogic) createUserJwtToken(userId string) (string, error) {
	// 将字符串类型的 UserId 转换为 int64 类型
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return "", err
	}

	// 创建jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userIdInt)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

// 处理不存在的用户
func (l *UserLogic) handleNonexistentUser(twitterId string, inviteId string, userAvatar string, userName string, twitterName string) (*types.UserLoginReply, error) {
	fmt.Println("handleNonexistentUser")

	fmt.Println("userAvatar: ", userAvatar)
	// 创建用户
	createUserResp, err := l.svcCtx.UserRpcClient.CreateUser(l.ctx, &userclient.CreateUserRequest{TwitterId: twitterId})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("createUserResp ")
	// 创建邀请
	if inviteId != "" {
		_, err = l.svcCtx.UserRpcClient.CreateInvite(l.ctx, &userclient.CreateInviteRequest{
			UserId:   createUserResp.UserId,
			InviteId: inviteId,
		})
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(createUserResp.UserId)

	// 添加信息
	_, err = l.svcCtx.UserRpcClient.AddUserInfo(l.ctx, &userclient.AddUserInfoRequest{
		UserId:      createUserResp.UserId,
		UserAvatar:  userAvatar,
		UserName:    userName,
		TwitterName: twitterName,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 创建jwt token
	jwtToken, err := l.createUserJwtToken(createUserResp.UserId)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginReply{
		Token: jwtToken,
	}, nil
}

// 处理已存在的用户
func (l *UserLogic) handleExistingUser(userId string) (*types.UserLoginReply, error) {
	jwtToken, err := l.createUserJwtToken(userId)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginReply{
		Token: jwtToken,
	}, nil
}

func (l *UserLogic) User(req *types.UserLogin) (resp *types.UserLoginReply, err error) {
	url := "http://47.243.34.199:9696/oauth/callback"
	method := "POST"

	// Construct the payload using the retrieved values
	payload := strings.NewReader(fmt.Sprintf(`{
    "code": "%s",
    "code_verifier": "%s"
	}`, req.Code, req.CodeVerifier))

	client := &http.Client{}
	reqs, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	reqs.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	reqs.Header.Add("Content-Type", "application/json")
	ress, err := client.Do(reqs)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ress.Body.Close()

	body, err := ioutil.ReadAll(ress.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	// Parse the response body into a ResponseData struct
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Access the parsed data
	fmt.Println("Access Token:", responseData.Data.AccessToken)
	fmt.Println("Name:", responseData.Data.Name)
	fmt.Println("User ID:", responseData.Data.UserID)
	fmt.Println("Username:", responseData.Data.Username)
	fmt.Println("Avatar:", responseData.Data.Avatar)
	// Access other fields as needed

	userResp, err := l.svcCtx.UserRpcClient.CheckTwitterId(l.ctx, &userclient.CheckTwitterIdRequest{TwitterId: responseData.Data.UserID})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return l.handleNonexistentUser(responseData.Data.UserID, req.InviteId, responseData.Data.Avatar, responseData.Data.Name, responseData.Data.Username) // 处理不存在的用户
		}
		return nil, err
	}

	// 添加在线用户
	_, err = l.svcCtx.Redis.Sadd("online_users", userResp.UserId)
	if err != nil {
		return nil, err
	}
	// 设置过期时间 一小时
	l.svcCtx.Redis.Expire("online_users", 3600)

	return l.handleExistingUser(userResp.UserId) // 处理已存在的用户
}
