package svc

import (
	"gateway/client/user"
	"gateway/internal/config"
	"google.golang.org/grpc"
	"log"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建 gRPC 连接
	conn, err := grpc.Dial(c.UserRpcConf, grpc.WithInsecure())
	if err != nil {
		// 处理错误情况
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}

	// 创建用户RPC客户端
	userRpcClient := user.NewUserClient(conn)

	return &ServiceContext{
		Config:        c,
		UserRpcClient: userRpcClient,
	}
}
