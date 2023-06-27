package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VoluntarilyTaskScheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVoluntarilyTaskScheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VoluntarilyTaskScheduleLogic {
	return &VoluntarilyTaskScheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VoluntarilyTaskScheduleLogic) VoluntarilyTaskSchedule(req *types.VoluntarilyTaskScheduleInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.VoluntarilyTaskSchedule(l.ctx, &taskclient.VoluntarilyTaskScheduleInput{
		UserId: req.UserId,
		TaskId: req.TaskId,
		Genre:  req.Genre,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
