package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskCallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskCallLogic {
	return &TaskCallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TaskCall 任务call
func (l *TaskCallLogic) TaskCall(req *types.TaskCallInput) (resp *types.TaskIDInquireInput, err error) {
	data, err := l.svcCtx.TaskClient.TaskCall(l.ctx, &taskclient.TaskCallInput{
		UserId: req.UserId,
		Sharer: req.Sharer,
		TaskID: req.TaskID,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.TaskIDInquireInput{
		TaskId: data.TaskId,
	}
	return
}
