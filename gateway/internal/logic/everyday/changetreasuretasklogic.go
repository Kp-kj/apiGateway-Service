package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeTreasureTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeTreasureTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeTreasureTaskLogic {
	return &ChangeTreasureTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeTreasureTaskLogic) ChangeTreasureTask(req *types.TreasureTaskInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
