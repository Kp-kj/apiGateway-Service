package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTreasureTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTreasureTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTreasureTaskListLogic {
	return &QueryTreasureTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTreasureTaskListLogic) QueryTreasureTaskList(req *types.TreasureTaskListInput) (resp *types.ReTreasureTaskSrt, err error) {
	// todo: add your logic here and delete this line

	return
}
