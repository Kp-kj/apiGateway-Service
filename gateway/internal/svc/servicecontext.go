package svc

import (
	"gateway/blockclient"
	"gateway/internal/config"
	"gateway/internal/middleware"
	"gateway/taskclient"
	"gateway/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	BlackMiddleware rest.Middleware
	UserRpcClient   userclient.User
	TaskClient      taskclient.Task
	BlockClient     blockclient.Block
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		BlackMiddleware: middleware.NewBlackMiddleware().Handle, // 黑名单中间件 初始化
		UserRpcClient:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		TaskClient:      taskclient.NewTask(zrpc.MustNewClient(c.TaskRpcConf)),
		BlockClient:     blockclient.NewBlock(zrpc.MustNewClient(c.BlockRpcConf)),
	}
}
