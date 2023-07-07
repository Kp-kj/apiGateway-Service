package adminmall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateActivityLogic) CreateActivity(req *types.CreateActivityInput) (resp *types.IsSuccessReply, err error) {

	result, err := l.svcCtx.BlockClient.CreateActivity(l.ctx, &blockclient.CreateActivityRequest{
		AdminuserId:       req.AdminUserID,
		CryptominerTypeid: req.CryptominerTypeID,
		UserAmount:        req.UserAmount,
		MinPrice:          float32(req.MinPrice),
		FirstBargainPer:   float32(req.FirstBargainPer),
		FriendBargainPer:  float32(req.FriendBargainPer),
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
