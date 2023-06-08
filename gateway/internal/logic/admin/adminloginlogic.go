package admin

import (
	"context"
	"gateway/userclient"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 创建jwt token 逻辑
func (l *AdminLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLogin) (resp *types.AdminLoginReply, err error) {
	userResp, err := l.svcCtx.UserRpcClient.AdminLogin(l.ctx, &userclient.AdminLoginRequest{
		AdminName:   req.AdminName,
		AdminPasswd: req.AdminPassword,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 创建jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.AdminAuth.AccessSecret, now, accessExpire, userResp.AdminUserId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.AdminLoginReply{
		Token: jwtToken,
	}, nil
}
