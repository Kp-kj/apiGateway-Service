package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTreasureTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTreasureTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTreasureTaskListLogic {
	return &GetTreasureTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTreasureTaskListLogic) GetTreasureTaskList(req *types.TreasureTaskListInput) (resp *types.ReTreasureTaskSrt, err error) {
	// todo: add your logic here and delete this line

	return
}
