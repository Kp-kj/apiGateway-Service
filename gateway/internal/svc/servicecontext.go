package svc

import (
	"gateway/internal/config"
	"gateway/internal/middleware"
	"gateway/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	BlackMiddleware rest.Middleware
	UserRpcClient   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		BlackMiddleware: middleware.NewBlackMiddleware().Handle, // 黑名单中间件 初始化
		UserRpcClient:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
