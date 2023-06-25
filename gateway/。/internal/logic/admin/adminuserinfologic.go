package admin

import (
	"context"

	"gateway/。/internal/svc"
	"gateway/。/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserInfoLogic {
	return &AdminUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUserInfoLogic) AdminUserInfo() (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
