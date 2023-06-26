package ping

import (
	"context"
	"fmt"
	"gateway/internal/svc"
	"gateway/userclient"
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
	// ping 用户服务测试
	pingResp, err := l.svcCtx.UserRpcClient.Ping(l.ctx, &userclient.Request{Ping: "ping"})
	fmt.Println("pingResp: ", pingResp)
	logx.Error("pingResp: ", pingResp) //测试用
	if err != nil {
		return err
	}
	return nil
}
