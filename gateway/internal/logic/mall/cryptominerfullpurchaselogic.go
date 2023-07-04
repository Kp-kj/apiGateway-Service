package mall

import (
	"context"
	"gateway/blockclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CryptominerFullPurchaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCryptominerFullPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CryptominerFullPurchaseLogic {
	return &CryptominerFullPurchaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CryptominerFullPurchaseLogic) CryptominerFullPurchase(req *types.CryptominerPurchaseInput) (resp *types.IsSuccessReply, err error) {
	result, err := l.svcCtx.BlockClient.CryptominerFullPurchase(l.ctx, &blockclient.CryptominerPurchaseRequest{
		UserId:        req.UserID,
		CryptominerId: req.CryptominerID,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
