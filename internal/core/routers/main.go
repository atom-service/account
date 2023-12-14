package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/protect-we-network/server/internal/core/handler/client"
	"github.com/protect-we-network/server/internal/core/handler/node"
	"github.com/protect-we-network/server/internal/core/handler/post"
	"github.com/protect-we-network/server/internal/core/handler/user"
	"github.com/protect-we-network/server/internal/core/middler"
)

func Init(engine *gin.Engine) {

	// 根路由
	globalGroup := engine.Group("")
	// 用户登录才能访问的路由
	userGroup := engine.Group("", middler.Auth(false))
	// admin 管理员才能访问的路由
	adminGroup := engine.Group("admin", middler.Auth(true))

	{
		handle := user.User{}
		globalGroup.POST("login", handle.Login)
		globalGroup.POST("logout", handle.Logout)
		globalGroup.POST("register", handle.Register)
	}

	clientGroup := globalGroup.Group("client")
	{
		client := client.Client{}
		clientGroup.GET("config", client.Config)
		clientGroup.GET("user-config", client.UserConfig)
		clientGroup.GET("update-user-config", client.UpdateUserConfig)
	}

	clientAdminGroup := adminGroup.Group("client")
	{
		client := client.ClientManage{}
		clientAdminGroup.POST("update-config", client.UpdateConfig)
	}

	// post 路由
	postGroup := userGroup.Group("post")
	{
		post := post.Post{}
		postGroup.GET("list", post.List)
		postGroup.GET("detail", post.Detail)
	}

	// post 管理路由
	postManageGroup := adminGroup.Group("post")
	{
		post := post.PostManage{}
		postManageGroup.POST("/create", post.Create)
		postManageGroup.POST("/delete", post.Delete)
		postManageGroup.POST("/update", post.Update)
	}

	// node 路由
	nodeGroup := userGroup.Group("node")
	{
		node := node.Node{}
		nodeGroup.GET("/config", node.Config)
		nodeGroup.PUT("/report-flow", node.ReportFlow)
		nodeGroup.GET("/available-user", node.AvailableUser)
	}

	// node 管理路由
	nodeManageGroup := adminGroup.Group("node")
	{
		node := node.NodeManage{}
		nodeManageGroup.GET("/create", node.Create)
		nodeManageGroup.GET("/delete", node.Delete)
		nodeManageGroup.GET("/update", node.Update)
		nodeManageGroup.POST("/update-config", node.UpdateConfig)
	}

	//
}
