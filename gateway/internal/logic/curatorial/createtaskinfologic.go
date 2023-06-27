package curatorial

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskInfoLogic {
	return &CreateTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskInfoLogic) CreateTaskInfo(req *types.CreatePublishTaskInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
