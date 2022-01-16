package routes

/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 14:47:22
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-12 20:03:36
 * @Description  : 佛祖保佑,永无BUG
 */

import (
	controllers "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/app/http/controllers/api/v1/auth"
	"github.com/ZimoBoy/gohub/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组, 我们所有的 v1 版本的路由都存放在这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			// 使用手机号注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			// 使用邮箱注册
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码,需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			// 发送手机验证码
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			// 发送Email验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			// 登录控制器
			lgc := new(auth.LoginController)
			// 使用手机号,短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 密码登录
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新令牌
			authGroup.POST("/login/refresh-token", middlewares.AuthJTW(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			// 通过手机号重置密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			// 通过邮箱重置密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJTW(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJTW(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJTW(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJTW(), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJTW(), uc.UpdatePassword)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("", cgc.Index)
			cgcGroup.POST("", middlewares.AuthJTW(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJTW(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJTW(), cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.POST("", middlewares.AuthJTW(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJTW(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJTW(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
		}

		lsc := new(controllers.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("", lsc.Index)
		}
	}
}
