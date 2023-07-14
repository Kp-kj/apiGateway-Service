package mall

import (
	"context"
	"encoding/json"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BargainPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBargainPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BargainPayLogic {
	return &BargainPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BargainPayLogic) BargainPay(req *types.BargainPayInput) (resp *types.IsSuccessReply, err error) {
	userId := l.ctx.Value("userId")
	var userID int64
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}
	result, err := l.svcCtx.BlockClient.BargainPay(l.ctx, &blockclient.BargainPayRequest{
		UserId:        userID,
		CryptominerId: req.CryptominerID,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
