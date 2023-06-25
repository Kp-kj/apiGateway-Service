package svc

import (
	"gateway/。/internal/config"
	"gateway/。/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	BlackMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		BlackMiddleware: middleware.NewBlackMiddleware().Handle,
	}
}
