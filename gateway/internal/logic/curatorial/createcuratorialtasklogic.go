package curatorial

import (
	"context"
	"fmt"
	"gateway/taskclient"
	"github.com/shopspring/decimal"
	"strings"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCuratorialTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCuratorialTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCuratorialTaskLogic {
	return &CreateCuratorialTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateCuratorialTask 创建策展任务
func (l *CreateCuratorialTaskLogic) CreateCuratorialTask(req *types.CreatePublishTaskInput) (resp *types.Mistake, err error) {
	// 判断推特文章地址是否正确
	if !strings.Contains(req.TweetAddress, "twitter.com/") {
		return &types.Mistake{Msg: "推特文章地址不正确"}, fmt.Errorf("推特文章地址不正确")
	} // 判断金额是否合规（是否小于0.01）
	if decimal.NewFromFloat(req.AwardAmount).LessThan(decimal.NewFromFloat(0.01)) {
		return &types.Mistake{Msg: "奖励金额最小为0"}, fmt.Errorf("奖励金额最小为0.01")
	}
	// 判断用户数是否合规
	if req.MaxUser < 10 {
		return &types.Mistake{Msg: "奖励用户数量不能少于10人"}, fmt.Errorf("奖励用户数量不能少于10人")
	}
	var taskDemand []*taskclient.TaskDemand
	for _, item := range req.TaskDemand {
		taskDemand = append(taskDemand, &taskclient.TaskDemand{
			TaskId:     uint64(item.TaskID),
			TaskName:   int32(item.TaskName),
			TaskDemand: item.TaskDemand,
			Article:    item.Article,
		})
	}
	err1, err := l.svcCtx.TaskClient.CreateCuratorialTask(l.ctx, &taskclient.CreatePublishTaskInput{
		Creator:      req.Creator,
		TweetAddress: req.TweetAddress,
		Label:        req.Label,
		AwardBudget:  req.AwardBudget,
		MaxUser:      req.MaxUser,
		AwardAmount:  req.AwardAmount,
		TaskBak:      taskDemand,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
