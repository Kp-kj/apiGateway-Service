package adminmall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartGoodLogic {
	return &StartGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartGoodLogic) StartGood(req *types.StartGoodInput) (resp *types.IsSuccessReply, err error) {
	result, err := l.svcCtx.BlockClient.StartGood(l.ctx, &blockclient.StartGoodRequest{
		GoodTypeid: req.GoodTypeID,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
