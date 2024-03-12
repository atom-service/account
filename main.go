package main

import (
	"net"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/internal/server"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/config"
	"google.golang.org/grpc"
)

func init() {
	config.Declare("port", "80", true, "服务监听的端口")
}

func main() {
	// 声明&初始化配置
	config.MustLoad()
	// 初始化数据库
	db.Init()

	listen, err := net.Listen("tcp", config.MustGet("port"))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	protos.RegisterSecretServiceServer(grpcServer, server.NewSecretServer())
	protos.RegisterAccountServiceServer(grpcServer, server.NewAccountServer())
	protos.RegisterPermissionServiceServer(grpcServer, server.NewPermissionServer())
	panic(grpcServer.Serve(listen))
}
