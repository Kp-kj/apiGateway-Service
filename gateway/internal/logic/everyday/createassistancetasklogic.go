package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAssistanceTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAssistanceTaskLogic {
	return &CreateAssistanceTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateAssistanceTask 创建用户发布助力任务表
func (l *CreateAssistanceTaskLogic) CreateAssistanceTask(req *types.UserPublishingAssistanceTaskInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.CreateAssistanceTask(l.ctx, &taskclient.CreateUserPublishingAssistanceTaskInput{
		UserId:   req.UserId,
		UserName: req.UserName,
		Avatar:   req.Avatar,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
