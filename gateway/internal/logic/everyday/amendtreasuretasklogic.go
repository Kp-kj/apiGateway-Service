package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendTreasureTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAmendTreasureTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendTreasureTaskLogic {
	return &AmendTreasureTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AmendTreasureTaskLogic) AmendTreasureTask(req *types.TreasureTaskSrtInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
