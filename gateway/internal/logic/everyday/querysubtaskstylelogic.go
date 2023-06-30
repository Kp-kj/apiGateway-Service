package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubtaskStyleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubtaskStyleLogic {
	return &QuerySubtaskStyleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubtaskStyleLogic) QuerySubtaskStyle(req *types.TaskIDInquireInput) (resp *types.ReSubtaskStyle, err error) {
	data, err := l.svcCtx.TaskClient.QuerySubtaskStyle(l.ctx, &taskclient.TaskIDInquireInput{Id: req.Id})
	var subtaskStyle []*types.SubtaskStyle
	for _, item := range data.SubtaskStyle {
		subtaskStyle = append(subtaskStyle, &types.SubtaskStyle{
			TaskId:         item.TaskId,
			TaskName:       item.TaskName,
			TaskNameEng:    item.TaskNameEng,
			TaskDetails:    item.TaskDetails,
			TaskDetailsEng: item.TaskDetailsEng,
			TaskStatus:     item.TaskStatus,
		})
	}
	return &types.ReSubtaskStyle{SubtaskStyle: subtaskStyle}, err
}
