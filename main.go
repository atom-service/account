package main

import (
	"net"

	"github.com/grpcbrick/account/provider"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/easysql"
	"github.com/yinxulai/goutils/grpc/interceptor"
	"google.golang.org/grpc"
)

func init() {
	config.SetStandard("mongo-url", "", true, "RPC 使用的 MONGODB 数据库配置")
	config.SetStandard("mysql-url", "", true, "RPC 使用的 MYSQL 数据库配置")
	config.SetStandard("rpc-port", ":3000", true, "RPC 服务监听的端口")
	config.CreateJSONTemplate("./config.template.json")
	config.LoadFlag()
}

func main() {
	var err error
	easysql.Init("mysql", config.MustGet("mysql-url"))
	lis, err := net.Listen("tcp", config.MustGet("rpc-port"))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(interceptor.NewCalllogs()...)
	standard.RegisterAccountServer(grpcServer, provider.NewService())
	panic(grpcServer.Serve(lis))
}
