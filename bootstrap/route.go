/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 14:42:18
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 14:42:27
 * @Description  : 佛祖保佑,永无BUG
 */

// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/ZimoBoy/gohub/app/http/middlewares"
	"net/http"
	"strings"

	"github.com/ZimoBoy/gohub/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是html的话
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			// 默认返回 json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义, 请确认 url 和请求方法是否正确",
			})
		}
	})
}
