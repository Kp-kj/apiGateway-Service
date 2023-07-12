package curatorial

import (
	"context"
	"gateway/taskclient"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTaskListLogic {
	return &QueryTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTaskList 查询任务列表
func (l *QueryTaskListLogic) QueryTaskList(req *types.PublishTaskInput) (resp *types.RePublishTask, err error) {
	data, err := l.svcCtx.TaskClient.QueryTaskList(l.ctx, &taskclient.PublishTaskInput{
		Status:   req.Status,
		CurrPage: req.CurrPage,
		MaxNum:   req.MaxNum})
	var rePublishTaskBak []*types.RePublishTaskBak
	for _, item := range data.RePublishTaskBak {
		// 获取用户信息
		user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: item.Creator})
		if err != nil {
			return nil, err
		}
		rePublishTaskBak = append(rePublishTaskBak, &types.RePublishTaskBak{
			TaskID:        item.TaskId,
			CreatedAt:     item.CreatedAt,
			Creator:       item.Creator,
			CreatorName:   user.UserName,
			CreatorNick:   user.TwitterName,
			CreatorAvatar: user.UserAvatar,
			Status:        item.Status,
			TweetDetails:  item.TweetDetails,
			TweetPicture:  item.TweetPicture,
			Label:         item.Label,
			AwardBudget:   item.AwardBudget,
			MaxUser:       item.MaxUser,
			AwardAmount:   item.AwardAmount,
			EndTime:       item.EndTime,
			Accomplish:    item.Accomplish,
		})
	}
	return &types.RePublishTask{
		PaginationData: types.PaginationData{
			Total:   data.PaginationData.Total,
			Page:    data.PaginationData.Page,
			PerPage: data.PaginationData.PerPage,
		},
		RePublishTaskBak: rePublishTaskBak,
	}, err
}
