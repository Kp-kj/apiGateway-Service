package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserPowerTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserPowerTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserPowerTaskLogic {
	return &CreateUserPowerTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateUserPowerTask 创建用户助力信息
func (l *CreateUserPowerTaskLogic) CreateUserPowerTask(req *types.CreateUserPowerTaskInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.CreateUserPowerTask(l.ctx, &taskclient.CreateUserPowerTaskInput{
		PublishesUserId: req.PublishesUserId,
		HelperUserId:    req.HelperUserId,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
