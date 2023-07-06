package mall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeGoodsPurchaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJudgeGoodsPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeGoodsPurchaseLogic {
	return &JudgeGoodsPurchaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JudgeGoodsPurchaseLogic) JudgeGoodsPurchase(req *types.JudgeGoodsPurchaseInput) (resp *types.JudgeGoodsPurchaseReply, err error) {

	result, err := l.svcCtx.BlockClient.JudgeGoodsPurchase(l.ctx, &blockclient.JudgeGoodsPurchaseRequest{
		UserId:   req.UserID,
		GoodName: req.GoodName,
	})

	return &types.JudgeGoodsPurchaseReply{IsPurchase: result.IsPurchase}, err

}
