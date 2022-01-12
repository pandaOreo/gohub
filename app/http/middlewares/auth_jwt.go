/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-12 20:10:23
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-12 20:10:54
 * @Description  : 佛祖保佑,永无BUG
 */

// Package middlewares Gin 中间件
package middlewares

import (
	"fmt"
	"github.com/ZimoBoy/gohub/pkg/config"
	"github.com/ZimoBoy/gohub/pkg/jwt"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJTW() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxxxx 中获取信息,并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParseToken(c)

		// JWT 解析失败, 有错误信息
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功,设置用户信息
		_user := user.Get(claims.UserID)
		if _user.Id == 0 {
			response.Unauthorized(c, "找不到对应用户,用户可能已删除")
			return
		}

		// 将用户信息存入 gin.Context 里,后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", _user.GetStringID())
		c.Set("current_user_name", _user.Name)
		c.Set("current_user", _user)

		c.Next()
	}
}
