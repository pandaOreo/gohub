/**
 * @Author: fanjinghua
 * @Description:
 * @File: password_controller
 * @Version: 1.0.0
 * @Date: 2022/1/12 21:13
 */

// Package auth 处理用户注册,登录,密码重置
package auth

import (
	v1 "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/app/requests"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseApiController
}

// ResetByPhone 使用手机号重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	// 2. 更新密码
	_user := user.GetByPhone(request.Phone)
	if _user.ID == 0 {
		response.Abort404(c)
	} else {
		_user.Password = request.Password
		_user.Save()
		response.Success(c)
	}
}

// ResetByEmail 使用邮箱重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	// 2. 更新密码
	_user := user.GetByEmail(request.Email)
	if _user.ID == 0 {
		response.Abort404(c)
	} else {
		_user.Password = request.Password
		_user.Save()
		response.Success(c)
	}
}
