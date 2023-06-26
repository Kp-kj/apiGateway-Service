package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLaunchTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserLaunchTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLaunchTaskListLogic {
	return &QueryUserLaunchTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserLaunchTaskList 查询个人发起任务列表+参与任务
func (l *QueryUserLaunchTaskListLogic) QueryUserLaunchTaskList(req *types.UserLaunchTaskListInput) (resp *types.RePublishTask, err error) {
	data, err := l.svcCtx.TaskClient.QueryUserLaunchTaskList(l.ctx, &taskclient.UserLaunchTaskListInput{
		UserId:   req.UserId,
		CurrPage: req.CurrPage,
		MaxNum:   req.MaxNum,
		Status:   req.Status,
	})
	var rePublishTaskBak []*types.RePublishTaskBak
	for _, item := range data.RePublishTaskBak {
		var taskDemand []*types.TaskDemand
		for _, item1 := range item.TaskDemand {
			taskDemand = append(taskDemand, &types.TaskDemand{
				TaskID:     uint(item1.TaskId),
				TaskName:   int(item1.TaskName),
				TaskDemand: item1.TaskDemand,
				Article:    item1.Article,
			})
		}
		rePublishTaskBak = append(rePublishTaskBak, &types.RePublishTaskBak{
			TaskID:        item.TaskId,
			CreatedAt:     item.CreatedAt,
			Creator:       item.Creator,
			CreatorName:   item.CreatorName,
			CreatorNick:   item.CreatorNick,
			CreatorAvatar: item.CreatorAvatar,
			Status:        item.Status,
			TweetDetails:  item.TweetDetails,
			TweetPicture:  item.TweetPicture,
			Label:         item.Label,
			AwardBudget:   item.AwardBudget,
			MaxUser:       int32(item.MaxUser),
			AwardAmount:   item.AwardAmount,
			EndTime:       item.EndTime,
			Accomplish:    int32(item.Accomplish),
			TaskDemand:    taskDemand,
		})
	}

	return &types.RePublishTask{
		PaginationData: types.PaginationData{
			Total:   data.PaginationData.Total,
			Page:    data.PaginationData.Page,
			PerPage: data.PaginationData.PerPage,
		},
		RePublishTaskBak: rePublishTaskBak,
	}, nil
	return
}
