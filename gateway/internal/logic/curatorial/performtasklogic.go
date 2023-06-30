package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformTaskLogic {
	return &PerformTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PerformTask 判断是否完成策展任务
func (l *PerformTaskLogic) PerformTask(req *types.PerformTaskInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.PerformTask(l.ctx, &taskclient.PerformTaskInput{
		TaskId: req.TaskId,
		UserId: req.UserId,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
