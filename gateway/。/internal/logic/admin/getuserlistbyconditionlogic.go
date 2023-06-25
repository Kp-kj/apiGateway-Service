package admin

import (
	"context"

	"gateway/。/internal/svc"
	"gateway/。/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListByConditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListByConditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListByConditionLogic {
	return &GetUserListByConditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListByConditionLogic) GetUserListByCondition(req *types.GetUserListByCondition) (resp *types.GetUserListByConditionReply, err error) {
	// todo: add your logic here and delete this line

	return
}
