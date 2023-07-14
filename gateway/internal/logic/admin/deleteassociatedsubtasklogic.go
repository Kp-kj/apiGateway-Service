package admin

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAssociatedSubtaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAssociatedSubtaskLogic {
	return &DeleteAssociatedSubtaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAssociatedSubtask 删除管理子任务
func (l *DeleteAssociatedSubtaskLogic) DeleteAssociatedSubtask(req *types.TaskIDInquireInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.DeleteAssociatedSubtask(l.ctx, &taskclient.TaskIDInquireInput{TaskId: req.TaskId})
	if err != nil {
		return nil, err
	}
	return &types.Mistake{Msg: err1.Msg}, err
}
