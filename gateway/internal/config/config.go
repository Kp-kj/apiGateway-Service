package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf zrpc.RpcClientConf
	Auth        struct {
		AccessSecret string
		AccessExpire int64
	}
	AdminAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
