package main

import (
	"net"

	"github.com/grpcbrick/account/provider"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/grpc/interceptor"
	"google.golang.org/grpc"
)

func init() {
	config.SetStandard("mysql", "", true, "RPC 使用的 MYSQL 数据库配置")
	config.SetStandard("port", ":3000", true, "RPC 服务监听的端口")
	config.CreateJSONTemplate("./config.template.json")
	config.LoadJSONFile("config.json")
	config.AutoLoad()
}

func main() {
	var err error
	provider.InitDB()

	rpcListenAddress, err := config.Get("port")
	lis, err := net.Listen("tcp", rpcListenAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(interceptor.NewCalllogs()...)
	standard.RegisterAccountServer(grpcServer, provider.NewService())
	panic(grpcServer.Serve(lis))
}
