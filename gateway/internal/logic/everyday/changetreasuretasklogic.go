package everyday

import (
	"context"
	"gateway/taskclient"

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

// ChangeTreasureTask 上架+删除宝箱样式
func (l *ChangeTreasureTaskLogic) ChangeTreasureTask(req *types.TreasureTaskInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.ChangeTreasureTask(l.ctx, &taskclient.TreasureTaskInput{
		Id:     req.Id,
		Status: req.Status,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
