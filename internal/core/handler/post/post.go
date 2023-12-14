package post

import "github.com/gin-gonic/gin"

type Post struct {
}

func (*Post) List(c *gin.Context)   {}
func (*Post) Detail(c *gin.Context) {}

type PostManage struct{}

func (*PostManage) Create(c *gin.Context) {}
func (*PostManage) Delete(c *gin.Context) {}
func (*PostManage) Update(c *gin.Context) {}
