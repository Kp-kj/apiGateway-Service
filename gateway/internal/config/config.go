package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf  zrpc.RpcClientConf
	TaskRpcConf  zrpc.RpcClientConf
	BlockRpcConf zrpc.RpcClientConf
	Auth         struct {
		AccessSecret string
		AccessExpire int64
	}
	AdminAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Address string
		Pass    string
	}
}
