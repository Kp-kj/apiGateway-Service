package everyday

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/taskclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendAssociatedSubtaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAmendAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendAssociatedSubtaskLogic {
	return &AmendAssociatedSubtaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AmendAssociatedSubtaskLogic) AmendAssociatedSubtask(req *types.AssociatedSubtaskSrt) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.AmendAssociatedSubtask(l.ctx, &taskclient.AssociatedSubtaskSrt{
		AssociatedId:   req.AssociatedId,
		TaskId:         req.TaskId,
		TreasureId:     req.TreasureId,
		TaskName:       req.TaskName,
		TaskNameEng:    req.TaskNameEng,
		TaskDetails:    req.TaskDetails,
		TaskDetailsEng: req.TaskDetailsEng,
		TaskStatus:     req.TaskStatus,
		Reward:         req.Reward,
		Experience:     req.Experience,
		Number:         req.Number,
		Article:        req.Article,
		Link:           req.Link,
		Label:          req.Label,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
