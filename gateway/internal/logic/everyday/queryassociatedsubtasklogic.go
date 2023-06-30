package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAssociatedSubtaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAssociatedSubtaskLogic {
	return &QueryAssociatedSubtaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryAssociatedSubtask 查询关联子任务
func (l *QueryAssociatedSubtaskLogic) QueryAssociatedSubtask(req *types.TaskIDInquireInput) (resp *types.ReAssociatedSubtask, err error) {
	data, err := l.svcCtx.TaskClient.QueryAssociatedSubtask(l.ctx, &taskclient.TaskIDInquireInput{Id: req.Id})
	var associatedSubtask []*types.AssociatedSubtask
	for _, item := range data.AssociatedSubtask {
		associatedSubtask = append(associatedSubtask, &types.AssociatedSubtask{
			TaskId:         item.TaskId,
			TaskName:       item.TaskName,
			TaskNameEng:    item.TaskNameEng,
			TaskDetails:    item.TaskDetails,
			TaskDetailsEng: item.TaskDetailsEng,
			TaskStatus:     item.TaskStatus,
			Reward:         item.Reward,
			Experience:     item.Experience,
			Number:         item.Number,
			Article:        item.Article,
			Link:           item.Link,
			Label:          item.Label,
			TreasureId:     item.TreasureId,
		})
	}
	return &types.ReAssociatedSubtask{AssociatedSubtask: associatedSubtask}, err
}
