package admin

import (
	"context"
	"gateway/taskclient"

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

// CreateSubtaskStyle 创建子任务样式
func (l *CreateSubtaskStyleLogic) CreateSubtaskStyle(req *types.UserIDInquireInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.CreateSubtaskStyle(l.ctx, &taskclient.UserIDInquireInput{UserId: req.UserId})
	return &types.Mistake{Msg: err1.Msg}, err
}
