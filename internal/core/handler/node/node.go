package node

import (
	"github.com/gin-gonic/gin"
	"github.com/protect-we-network/server/internal/packages/logger"
	"github.com/protect-we-network/server/internal/packages/node"
)

type Node struct {
}

type ReportFlowRequest struct {
	Data []*node.FlowReport
}

// 节点上报信息
func (*Node) ReportFlow(c *gin.Context) {
	var req ReportFlowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		logger.Error(err)
		return
	}

	// 最大保存时间
	err := node.UpdateFlowReport(c.Request.Context(), req.Data)
	if err != nil {
		c.JSON(500, gin.H{"error": "上报流量出错"})
		logger.Error(err)
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

// 节点获取配置
func (*Node) Config(c *gin.Context) {

}

// 节点获取可用用户列表
func (*Node) AvailableUser(c *gin.Context) {
}

type NodeManage struct{}

func (*NodeManage) Create(c *gin.Context)       {}
func (*NodeManage) Delete(c *gin.Context)       {}
func (*NodeManage) Update(c *gin.Context)       {}
func (*NodeManage) UpdateConfig(c *gin.Context) {}
