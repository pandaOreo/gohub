/**
 * @Author: fanjinghua
 * @Description:
 * @File: signup_controller
 * @Version: 1.0.0
 * @Date: 2022/1/8 20:32
 */

// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	v1 "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/app/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败, 返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了,中断请求;
		return
	}

	//表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零即通过,大于0即有错误发生
	if len(errs) > 0 {
		// 验证失败, 返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
