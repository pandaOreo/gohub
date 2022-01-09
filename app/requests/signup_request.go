/**
 * @Author: fanjinghua
 * @Description:
 * @File: signup_request
 * @Version: 1.0.0
 * @Date: 2022/1/9 09:24
 */

// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidatePhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项,参数与名称为phone",
			"digits:手机号长度必须为11位的数字",
		},
	}

	return validate(data, rules, messages)
}

func ValidateEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于4",
			"max:Email 长度需小于30",
			"email:Email 格式不正确,请 提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}
