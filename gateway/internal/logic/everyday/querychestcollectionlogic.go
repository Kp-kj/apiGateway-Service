package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryChestCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryChestCollectionLogic {
	return &QueryChestCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryChestCollection 获取用户宝箱进度
func (l *QueryChestCollectionLogic) QueryChestCollection(req *types.UserIDInquireInput) (resp *types.ReChestCollectionSrt, err error) {
	data, err := l.svcCtx.TaskClient.QueryChestCollection(l.ctx, &taskclient.UserIDInquireInput{UserId: req.UserId})
	var treasureTaskStageSeed []*types.TreasureTaskStageSeed
	for _, item := range data.TreasureTaskStage {
		treasureTaskStageSeed = append(treasureTaskStageSeed, &types.TreasureTaskStageSeed{
			Treasure:         item.Treasure,
			TreasureSequence: item.TreasureSequence,
			StageExperience:  item.StageExperience,
			TreasureRatio:    item.TreasureRatio,
			StageReward:      item.StageReward,
		})
	}
	var associatedSubtaskSeed []*types.AssociatedSubtaskSeed
	for _, item1 := range data.AssociatedSubtask {
		associatedSubtaskSeed = append(associatedSubtaskSeed, &types.AssociatedSubtaskSeed{
			TreasureId:     item1.TreasureId,
			TaskId:         item1.TaskId,
			TaskName:       item1.TaskName,
			TaskNameEng:    item1.TaskNameEng,
			TaskDetails:    item1.TaskDetails,
			TaskDetailsEng: item1.TaskDetailsEng,
			TaskStatus:     item1.TaskStatus,
			Reward:         item1.Reward,
			Experience:     item1.Experience,
			Number:         item1.Number,
			Article:        item1.Article,
			Link:           item1.Link,
			Label:          item1.Label,
			Complete:       item1.Complete,
		})
	}
	return &types.ReChestCollectionSrt{
		SerId:                 data.SerId,
		DemandIntegral:        data.DemandIntegral,
		ChestAmount:           data.ChestAmount,
		RewardQuantity:        data.RewardQuantity,
		TreasureTaskStageSeed: treasureTaskStageSeed,
		AssociatedSubtaskSeed: associatedSubtaskSeed,
	}, err
}
