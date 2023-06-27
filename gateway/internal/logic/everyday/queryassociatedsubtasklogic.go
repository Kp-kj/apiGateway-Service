package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAssociatedSubtaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAssociatedSubtaskLogic {
	return &QueryAssociatedSubtaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAssociatedSubtaskLogic) QueryAssociatedSubtask(req *types.TaskIDInquireInput) (resp *types.ReAssociatedSubtask, err error) {
	// todo: add your logic here and delete this line

	return
}
