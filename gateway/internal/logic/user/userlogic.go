package user

import (
	"context"
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

// 处理不存在的用户
func (l *UserLogic) handleNonexistentUser(twitterId string, inviteId string) (*types.UserLoginReply, error) {
	// 创建用户
	createUserResp, err := l.svcCtx.UserRpcClient.CreateUser(l.ctx, &userclient.CreateUserRequest{TwitterId: twitterId})
	if err != nil {
		return nil, err
	}

	// 创建邀请
	_, err = l.svcCtx.UserRpcClient.CreateInvite(l.ctx, &userclient.CreateInviteRequest{
		UserId:   createUserResp.UserId,
		InviteId: inviteId,
	})
	if err != nil {
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
	// todo: 调用推特服务获取 twitterId 和 用户 创建

	twitterId := "12321312"

	userResp, err := l.svcCtx.UserRpcClient.CheckTwitterId(l.ctx, &userclient.CheckTwitterIdRequest{TwitterId: twitterId})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return l.handleNonexistentUser(twitterId, req.InviteId) // 处理不存在的用户
		}
		return nil, err
	}

	return l.handleExistingUser(userResp.UserId) // 处理已存在的用户
}
