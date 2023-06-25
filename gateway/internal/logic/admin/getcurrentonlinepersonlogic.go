package admin

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentOnlinePersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentOnlinePersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentOnlinePersonLogic {
	return &GetCurrentOnlinePersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetCurrentOnlinePerson 获取当前在线人数
func (l *GetCurrentOnlinePersonLogic) GetCurrentOnlinePerson() (resp *types.CurrentOnlinePersonReply, err error) {
	// todo: add your logic here and delete this line

	return &types.CurrentOnlinePersonReply{CurrentOnlinePerson: 0}, nil
}
