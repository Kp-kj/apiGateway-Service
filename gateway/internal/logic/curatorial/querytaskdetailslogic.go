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

// QueryTaskDetails 获取任务详情
func (l *QueryTaskDetailsLogic) QueryTaskDetails(req *types.TaskDetailsInput) (resp *types.ReTaskDetails, err error) {
	date, err := l.svcCtx.TaskClient.QueryTaskDetails(l.ctx, &taskclient.TaskDetailsInput{
		TaskId: req.TaskId,
		UserId: req.UserId,
	})
	// 获取用户信息
	userPublishTask, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: date.RePublishTaskSrt.Creator})
	if err != nil {
		return nil, err
	}
	rePublishTaskSrt := types.RePublishTaskBak{
		TaskID:        date.RePublishTaskSrt.TaskId,
		CreatedAt:     date.RePublishTaskSrt.CreatedAt,
		Creator:       date.RePublishTaskSrt.Creator,
		CreatorName:   userPublishTask.UserName,
		CreatorNick:   userPublishTask.TwitterName,
		CreatorAvatar: userPublishTask.UserAvatar,
		Status:        date.RePublishTaskSrt.Status,
		Label:         date.RePublishTaskSrt.Label,
		TweetDetails:  date.RePublishTaskSrt.TweetDetails,
		TweetPicture:  date.RePublishTaskSrt.TweetPicture,
		AwardBudget:   date.RePublishTaskSrt.AwardBudget,
		MaxUser:       date.RePublishTaskSrt.MaxUser,
		AwardAmount:   date.RePublishTaskSrt.AwardAmount,
		EndTime:       date.RePublishTaskSrt.EndTime,
		Accomplish:    date.RePublishTaskSrt.Accomplish,
		Call:          date.RePublishTaskSrt.Call,
	}
	var participant []*types.ParticipantBak
	for _, item := range date.Participant {
		// 获取用户信息
		user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: item.UserId})
		if err != nil {
			return nil, err
		}
		participant = append(participant, &types.ParticipantBak{
			Id:        item.Id,
			CreatedAt: item.CreatedAt,
			UserId:    item.UserId,
			TaskId:    item.TaskId,
			UserName:  user.UserName,
			NickName:  user.TwitterName,
			Avatar:    user.UserAvatar,
		})
	}
	return &types.ReTaskDetails{
		RePublishTaskBak: rePublishTaskSrt,
		ParticipantBak:   participant,
	}, err
}
