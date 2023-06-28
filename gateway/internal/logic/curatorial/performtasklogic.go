package curatorial

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerformTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformTaskLogic {
	return &PerformTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerformTaskLogic) PerformTask(req *types.PerformTaskInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
