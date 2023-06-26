package curatorial

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLaunchTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLaunchTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLaunchTaskListLogic {
	return &GetUserLaunchTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLaunchTaskListLogic) GetUserLaunchTaskList(req *types.UserLaunchTaskListInput) (resp *types.RePublishTask, err error) {

	return
}
