package mall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropPurchaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPropPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropPurchaseLogic {
	return &PropPurchaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PropPurchaseLogic) PropPurchase(req *types.PropPurchaseInput) (resp *types.IsSuccessReply, err error) {

	result, err := l.svcCtx.BlockClient.PropPurchase(l.ctx, &blockclient.PropPurchaseRequest{
		UserId:       req.UserID,
		PropId:       req.PropID,
		GoodQuantity: req.GoodQuantity,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
