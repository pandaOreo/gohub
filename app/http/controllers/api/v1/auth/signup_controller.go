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
	v1 "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/app/requests"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数,并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidatePhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 获取请求参数,并做表单验证
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
