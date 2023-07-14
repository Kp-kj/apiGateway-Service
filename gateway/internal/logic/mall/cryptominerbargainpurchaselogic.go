package mall

import (
	"context"
	"encoding/json"
	"gateway/blockclient"
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
	userId := l.ctx.Value("userId")
	var userID int64
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}

	result, err := l.svcCtx.BlockClient.CryptominerBargainPurchase(l.ctx, &blockclient.CryptominerBargainRequest{
		UserId:        userID,
		CryptominerId: req.CryptominerID,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.CryptominerBargainReply{BargainID: result.BargainId}, err
}
