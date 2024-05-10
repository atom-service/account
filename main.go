package main

import (
	"context"
	"strconv"
	"time"

	"github.com/atom-service/account/internal/config"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/internal/server"
)

func main() {
	// 声明&初始化配置
	config.MustInit()

	context, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	// 初始化 model
	if err := model.Init(context); err != nil {
		panic(err)
	}

	listenAddress := ":" + strconv.Itoa(config.Service.Port) 
	panic(server.StartServer(listenAddress))
}
