package curatorial

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/taskclient"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallSkipMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallSkipMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallSkipMessageLogic {
	return &CallSkipMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CallSkipMessage call跳转获取分享信息
func (l *CallSkipMessageLogic) CallSkipMessage(req *types.TaskCallInput) (resp *types.ReTaskDetails, err error) {
	userId := l.ctx.Value("userId")
	var userIdNum string
	if v, ok := userId.(json.Number); ok {
		userIdNum = v.String()
	}
	if req.UserId != userIdNum {
		return nil, fmt.Errorf("用户ID解析不正确")
	}
	date, err := l.svcCtx.TaskClient.CallSkipMessage(l.ctx, &taskclient.TaskCallInput{
		TaskID: req.TaskID,
		UserId: req.UserId,
		Sharer: req.Sharer,
	})
	if err != nil {
		return nil, err
	}
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
