package routes

import (
	"MyGinBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由入口文件

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() // 加了两个中间件（默认）

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)
}
