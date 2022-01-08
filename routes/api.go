package routes

/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 14:47:22
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 14:49:10
 * @Description  : 佛祖保佑,永无BUG
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组, 我们所有的 v1 版本的路由都存放在这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
