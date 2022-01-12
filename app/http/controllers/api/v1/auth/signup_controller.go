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
	"github.com/ZimoBoy/gohub/pkg/jwt"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExist 判断手机号是否存在
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

// IsEmailExist 判断邮箱是否存在
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

// SignupUsingPhone 使用手机号 + 验证码注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功,创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试~")
	}
}
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功,创建数据
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试~")
	}
}
