/**
 * @Author: fanjinghua
 * @Description:
 * @File: auth
 * @Version: 1.0.0
 * @Date: 2022/1/12 13:46
 */

// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"github.com/ZimoBoy/gohub/app/models/user"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}
