package routes

/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 14:47:22
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 14:49:10
 * @Description  : 佛祖保佑,永无BUG
 */

import (
	"github.com/ZimoBoy/gohub/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组, 我们所有的 v1 版本的路由都存放在这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
