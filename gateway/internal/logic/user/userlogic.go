package user

import (
	"context"
	"fmt"
	"gateway/userclient"
	"github.com/golang-jwt/jwt/v4"
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

func (l *UserLogic) User(req *types.UserLogin) (resp *types.UserLoginReply, err error) {
	// todo: 调用推特服务获取 twitterId 和 用户 创建
	//用推特url 发送给 推特服务拿到 twitterid和twittername
	//这部分省略直接已经拿到了
	twitterId := "1498198918672580608 "
	//twitterName := "Ed17899676"
	//验证是否有这个用户
	pingResp, err := l.svcCtx.UserRpcClient.CheckTwitterId(l.ctx, &userclient.CheckTwitterIdRequest{TwitterId: twitterId})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") { //处理rpc错误 不存在这个用户
			createUserResp, err := l.svcCtx.UserRpcClient.CreateUser(l.ctx, &userclient.CreateUserRequest{TwitterId: twitterId})
			if err != nil {
				return nil, err
			}

			jwtToken, err := l.createUserJwtToken(createUserResp.UserId) //创建jwt
			if err != nil {
				return nil, err
			}
			return &types.UserLoginReply{
				Token: jwtToken,
			}, nil
		}
		return nil, err
	}
	jwtToken, err := l.createUserJwtToken(pingResp.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println(jwtToken)
	return &types.UserLoginReply{
		Token: jwtToken,
	}, nil
}
