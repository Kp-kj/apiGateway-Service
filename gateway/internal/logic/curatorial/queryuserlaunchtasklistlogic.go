package curatorial

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/taskclient"
	"gateway/userclient"

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

// QueryUserLaunchTaskList 查询用户启动任务列表
func (l *QueryUserLaunchTaskListLogic) QueryUserLaunchTaskList(req *types.UserLaunchTaskListInput) (resp *types.UserLaunchTaskList, err error) {
	data, err := l.svcCtx.TaskClient.QueryUserLaunchTaskList(l.ctx, &taskclient.UserLaunchTaskListInput{
		UserId:   req.UserId,
		CurrPage: req.CurrPage,
		MaxNum:   req.MaxNum,
		Status:   req.Status,
	})
	var rePersonalList []*types.PersonalList
	for _, item := range data.PersonalList {
		rePersonalList = append(rePersonalList, &types.PersonalList{
			TaskId:       item.TaskId,
			CreatedAt:    item.CreatedAt,
			Creator:      item.Creator,
			Label:        item.Label,
			Status:       item.Status,
			TweetDetails: item.TweetDetails,
			TweetPicture: item.TweetPicture,
			AwardBudget:  item.AwardBudget,
			MaxUser:      item.MaxUser,
			Accomplish:   item.Accomplish,
		})
	}
	// 判断用户是否存在
	user, _ := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: req.UserId})
	if user == nil {
		return nil, err
	}
	return &types.UserLaunchTaskList{
		UserId:           req.UserId,
		UserName:         user.UserName,
		UserNick:         user.TwitterName,
		UserAvatar:       user.UserAvatar,
		LaunchAmount:     data.LaunchAmount,
		AccomplishAmount: data.AccomplishAmount,
		PaginationData: types.PaginationData{
			Total:   data.PaginationData.Total,
			Page:    data.PaginationData.Page,
			PerPage: data.PaginationData.PerPage,
		},
		PersonalList: rePersonalList,
	}, err
}
