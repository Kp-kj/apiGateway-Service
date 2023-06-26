package curatorial

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
