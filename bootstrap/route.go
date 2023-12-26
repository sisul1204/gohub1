// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(route *gin.Engine) {
	registerGlobalMiddleWare(route)
	setup404Handler(route)
	routes.RegisterAPIRoutes(route)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认url和请求方法是否正确",
			})
		}
	})
}
