package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubtaskStyleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubtaskStyleLogic {
	return &QuerySubtaskStyleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubtaskStyleLogic) QuerySubtaskStyle(req *types.TaskIDInquireInput) (resp *types.ReSubtaskStyle, err error) {
	// todo: add your logic here and delete this line

	return
}
