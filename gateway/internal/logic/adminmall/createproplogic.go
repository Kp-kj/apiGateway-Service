package adminmall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropLogic {
	return &CreatePropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePropLogic) CreateProp(req *types.CreatePropInput) (resp *types.IsSuccessReply, err error) {
	result, err := l.svcCtx.BlockClient.CreateProp(l.ctx, &blockclient.CreatePropRequest{
		AdminuserId:  req.AdminUserID,
		PropName:     req.PropName,
		PropPicture:  req.PropPicture,
		PropDescribe: req.PropDescribe,
		PropPrice:    req.PropPrice,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err

	return
}
