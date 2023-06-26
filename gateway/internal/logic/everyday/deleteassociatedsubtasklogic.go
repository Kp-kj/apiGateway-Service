package everyday

import (
	"context"

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

func (l *DeleteAssociatedSubtaskLogic) DeleteAssociatedSubtask(req *types.TaskIDInquireInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
