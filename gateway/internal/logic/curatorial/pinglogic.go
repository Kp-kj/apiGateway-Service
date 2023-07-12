package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.TaskIDInquireInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.Ping(l.ctx, &taskclient.TaskIDInquireInput{TaskId: req.TaskId})
	return &types.Mistake{Msg: err1.Msg}, err
}
