package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubtaskStyleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubtaskStyleLogic {
	return &CreateSubtaskStyleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSubtaskStyleLogic) CreateSubtaskStyle(req *types.UserIDInquireInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
