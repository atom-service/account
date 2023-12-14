package client

import (
	"github.com/gin-gonic/gin"
	"github.com/protect-we-network/server/internal/packages/client"
	"github.com/protect-we-network/server/internal/packages/http"
)

type Client struct{}

func (*Client) Config(c *gin.Context) {
	keys, _ := c.GetQueryArray("key")
	params := &client.QueryConfigParams{Keys: keys}
	configs, err := client.QueryConfig(c.Request.Context(), params)
	if err != nil {
		c.JSON(200, http.CreateHttpResult(500, nil, nil))
		return
	}

	c.JSON(200, http.CreateHttpResult(200, configs, nil))
}

func (*Client) UserConfig(c *gin.Context) {
	keys, _ := c.GetQueryArray("key")
	params := &client.QueryConfigParams{Keys: keys}
	configs, err := client.QueryConfig(c.Request.Context(), params)
	if err != nil {
		c.JSON(200, http.CreateHttpResult(500, nil, nil))
		return
	}

	c.JSON(200, http.CreateHttpResult(200, configs, nil))

}
func (*Client) UpdateUserConfig(c *gin.Context) {
	
}

type ClientManage struct{}

func (*ClientManage) UpdateConfig(c *gin.Context) {

}
