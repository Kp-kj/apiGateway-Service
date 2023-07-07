package adminmall

import (
	"context"
	"gateway/blockclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartActivityLogic {
	return &StartActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartActivityLogic) StartActivity(req *types.StartActivityInput) (resp *types.IsSuccessReply, err error) {
	result, err := l.svcCtx.BlockClient.StartActivity(l.ctx, &blockclient.StartActivityRequest{
		ActivityId: req.ActivityID,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
