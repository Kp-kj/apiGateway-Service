package admin

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

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

// AmendAssociatedSubtask 创建关联子任务+编辑关联子任务
func (l *AmendAssociatedSubtaskLogic) AmendAssociatedSubtask(req *types.AssociatedSubtaskSrt) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.AmendAssociatedSubtask(l.ctx, &taskclient.AssociatedSubtaskSrt{
		SubtaskId:      req.SubtaskId,
		TaskId:         req.TaskId,
		TreasureId:     req.TreasureId,
		TaskName:       req.TaskName,
		TaskNameEng:    req.TaskNameEng,
		TaskDetails:    req.TaskDetails,
		TaskDetailsEng: req.TaskDetailsEng,
		TaskStatus:     req.TaskStatus,
		Experience:     req.Experience,
		Number:         req.Number,
		Article:        req.Article,
		Link:           req.Link,
		Label:          req.Label,
	})
	if err != nil {
		return nil, err
	}
	return &types.Mistake{Msg: err1.Msg}, nil
}
