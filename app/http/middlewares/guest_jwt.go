/**
 * @Author: fanjinghua
 * @Description:
 * @File: guest_jwt
 * @Version: 1.0.0
 * @Date: 2022/1/12 21:03
 */

package middlewares

import (
	"github.com/ZimoBoy/gohub/pkg/jwt"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			// 解析 token 成功, 说明登录成功了
			_, err := jwt.NewJWT().ParseToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
