package ping

import (
	"context"
	"fmt"
	"gateway/client/user"
	"gateway/internal/svc"
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

func (l *PingLogic) Ping() error {
	pingResp, err := l.svcCtx.UserRpcClient.Ping(l.ctx, &user.Request{Ping: "ping"}, nil)
	if err != nil {
		return err
	}
	fmt.Println(pingResp.Pong)
	return nil
}
