package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//handle方法
//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("blog")
	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	//健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
