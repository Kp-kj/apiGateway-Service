package curatorial

import (
	"context"
	"gateway/taskclient"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTaskDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTaskDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTaskDetailsLogic {
	return &QueryTaskDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTaskDetailsLogic) QueryTaskDetails(req *types.TaskDetailsInput) (resp *types.ReTaskDetails, err error) {
	date, err := l.svcCtx.TaskClient.QueryTaskDetails(l.ctx, &taskclient.TaskDetailsInput{
		TaskId: req.TaskId,
		UserId: req.UserId,
	})
	var taskDemand []*types.TaskDemand
	for _, item := range date.RePublishTaskSrt.TaskDemand {
		taskDemand = append(taskDemand, &types.TaskDemand{
			TaskID:     uint(item.TaskId),
			TaskName:   int(item.TaskName),
			TaskDemand: item.TaskDemand,
			Article:    item.Article,
		})
	}
	//// 获取用户信息
	//userPublishTask, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: date.RePublishTaskSrt.UserId})
	//if err != nil {
	//	return nil, err
	//}
	rePublishTaskSrt := types.RePublishTaskBak{
		TaskID:        date.RePublishTaskSrt.TaskId,
		CreatedAt:     date.RePublishTaskSrt.CreatedAt,
		Creator:       date.RePublishTaskSrt.Creator,
		CreatorName:   date.RePublishTaskSrt.CreatorName,
		CreatorNick:   date.RePublishTaskSrt.CreatorNick,
		CreatorAvatar: date.RePublishTaskSrt.CreatorAvatar,
		Status:        date.RePublishTaskSrt.Status,
		TweetDetails:  date.RePublishTaskSrt.TweetDetails,
		TweetPicture:  date.RePublishTaskSrt.TweetPicture,
		Label:         date.RePublishTaskSrt.Label,
		AwardBudget:   date.RePublishTaskSrt.AwardBudget,
		MaxUser:       int32(date.RePublishTaskSrt.MaxUser),
		AwardAmount:   date.RePublishTaskSrt.AwardAmount,
		EndTime:       date.RePublishTaskSrt.EndTime,
		Accomplish:    int32(date.RePublishTaskSrt.Accomplish),
		TaskDemand:    taskDemand,
	}
	var participant []*types.ParticipantBak
	for _, item := range date.Participant {
		// 获取用户信息
		user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: item.UserId})
		if err != nil {
			return nil, err
		}
		participant = append(participant, &types.ParticipantBak{
			UserId:      item.UserId,
			UserName:    user.UserName,
			NickName:    user.TwitterName,
			Avatar:      user.UserAvatar,
			AwardAmount: item.AwardAmount,
			TaskID:      item.TaskID,
			Status:      item.Status,
		})
	}
	var taskDemandBak []*types.TaskDemandBak
	for _, item := range date.TaskDemandBak {
		taskDemandBak = append(taskDemandBak, &types.TaskDemandBak{
			TaskID:     item.TaskID,
			TaskName:   item.TaskName,
			TaskDemand: item.TaskDemand,
			Article:    item.Article,
			Status:     item.Done,
		})
	}
	return &types.ReTaskDetails{
		RePublishTaskBak: rePublishTaskSrt,
		ParticipantBak:   participant,
		TaskDemandBak:    taskDemandBak,
	}, err
}
