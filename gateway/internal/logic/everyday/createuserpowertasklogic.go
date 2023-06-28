package everyday

import (
	"context"

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

func (l *CreateUserPowerTaskLogic) CreateUserPowerTask(req *types.CreateUserPowerTaskInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
