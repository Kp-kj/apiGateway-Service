package admin

import (
	"context"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegisteredPopulationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegisteredPopulationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegisteredPopulationLogic {
	return &GetRegisteredPopulationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetRegisteredPopulation 获取注册人数
func (l *GetRegisteredPopulationLogic) GetRegisteredPopulation() (resp *types.RegisteredPopulationReply, err error) {
	// todo: add your logic here and delete this line
	RegisteredPopulation, err := l.svcCtx.UserRpcClient.RegisterCount(l.ctx, &userclient.Request{Ping: "ping"})
	if err != nil {
		return nil, err
	}
	return &types.RegisteredPopulationReply{RegisteredPopulation: RegisteredPopulation.Count}, nil
}
