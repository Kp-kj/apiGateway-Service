package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GatAssociatedSubtaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGatAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatAssociatedSubtaskLogic {
	return &GatAssociatedSubtaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GatAssociatedSubtaskLogic) GatAssociatedSubtask(req *types.TaskIDInquireInput) (resp *types.ReAssociatedSubtask, err error) {
	// todo: add your logic here and delete this line

	return
}
