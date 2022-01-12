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
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 使用手机号注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 使用邮箱注册
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码,需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送手机验证码
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 发送Email验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			// 登录控制器
			lgc := new(auth.LoginController)
			// 使用手机号,短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
		}
	}
}
