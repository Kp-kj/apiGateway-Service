package mall

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CryptominerBargainPurchaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCryptominerBargainPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CryptominerBargainPurchaseLogic {
	return &CryptominerBargainPurchaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CryptominerBargainPurchaseLogic) CryptominerBargainPurchase(req *types.CryptominerBargainInput) (resp *types.CryptominerBargainReply, err error) {
	// todo: add your logic here and delete this line

	return
}
