package everyday

import (
	"context"

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

func (l *AmendAssociatedSubtaskLogic) AmendAssociatedSubtask(req *types.AssociatedSubtaskSrt) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
