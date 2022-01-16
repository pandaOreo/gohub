/**
 * @Author: fanjinghua
 * @Description:
 * @File: force_ua
 * @Version: 1.0.0
 * @Date: 2022/1/16 19:34
 */

// Package middlewares Gin 中间件
package middlewares

import (
	"errors"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}
		c.Next()
	}
}
