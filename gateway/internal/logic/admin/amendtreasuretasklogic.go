package admin

import (
	"context"
	"fmt"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendTreasureTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAmendTreasureTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendTreasureTaskLogic {
	return &AmendTreasureTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AmendTreasureTask 创建宝箱样式+编辑宝箱样式
func (l *AmendTreasureTaskLogic) AmendTreasureTask(req *types.TreasureTaskSrtInput) (resp *types.Mistake, err error) {
	var treasureTaskStage []*taskclient.TreasureTaskStage
	for i, item := range req.TreasureTaskStage {
		if item.TreasureSequence != int64(i+1) {
			return &types.Mistake{Msg: "宝箱阶段不合规"}, nil
		}
		treasureTaskStage = append(treasureTaskStage, &taskclient.TreasureTaskStage{
			Treasure:         item.Treasure,
			TreasureSequence: item.TreasureSequence,
			StageExperience:  item.StageExperience,
			StageReward:      item.StageReward,
		})
	}
	treasureTask := &taskclient.TreasureTaskSrtInput{
		TreasureId:        req.TreasureId,
		Name:              req.Name,
		DemandIntegral:    req.DemandIntegral,
		ExperienceReward:  req.ExperienceReward,
		RewardQuantity:    req.RewardQuantity,
		TreasureTaskStage: treasureTaskStage,
	}
	err1, err := l.svcCtx.TaskClient.AmendTreasureTask(l.ctx, treasureTask)
	if err != nil {
		fmt.Printf("xc:%v\n", err)
		return nil, err
	}
	return &types.Mistake{Msg: err1.Msg}, nil
}
