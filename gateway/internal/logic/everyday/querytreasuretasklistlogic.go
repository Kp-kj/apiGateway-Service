package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTreasureTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTreasureTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTreasureTaskListLogic {
	return &QueryTreasureTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTreasureTaskListLogic) QueryTreasureTaskList(req *types.TreasureTaskListInput) (resp *types.ReTreasureTaskSrt, err error) {
	data, err := l.svcCtx.TaskClient.QueryTreasureTaskList(l.ctx, &taskclient.TreasureTaskListInput{
		Name:     req.Name,
		Reward:   req.Reward,
		CurrPage: req.CurrPage,
		MaxNum:   req.MaxNum,
	})
	var treasureTaskSrt []*types.TreasureTaskSrtInput
	for _, item := range data.TreasureTaskSrtInput {
		var treasureTaskStage []*types.TreasureTaskStage
		for _, item1 := range item.TreasureTaskStage {
			treasureTaskStage = append(treasureTaskStage, &types.TreasureTaskStage{
				Treasure:         item1.Treasure,
				TreasureSequence: item1.TreasureSequence,
				StageExperience:  item1.StageExperience,
				StageReward:      item1.StageReward,
			})
		}
		treasureTaskSrt = append(treasureTaskSrt, &types.TreasureTaskSrtInput{
			Id:                item.Id,
			Name:              item.Name,
			DemandIntegral:    item.DemandIntegral,
			TaskReward:        item.TaskReward,
			ExperienceReward:  item.ExperienceReward,
			RewardQuantity:    item.RewardQuantity,
			TreasureTaskStage: treasureTaskStage,
		})
	}

	return &types.ReTreasureTaskSrt{
		PaginationData: types.PaginationData{
			Total:   data.PaginationData.Total,
			Page:    data.PaginationData.Page,
			PerPage: data.PaginationData.PerPage,
		},
		TreasureTaskSrtInput: treasureTaskSrt,
	}, nil
}
