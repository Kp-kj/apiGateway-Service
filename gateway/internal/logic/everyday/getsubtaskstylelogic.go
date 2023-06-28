package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubtaskStyleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubtaskStyleLogic {
	return &GetSubtaskStyleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubtaskStyleLogic) GetSubtaskStyle(req *types.TaskIDInquireInput) (resp *types.ReSubtaskStyle, err error) {
	// todo: add your logic here and delete this line

	return
}
