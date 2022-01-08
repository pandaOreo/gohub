/**
 * @Author: fanjinghua
 * @Description:
 * @File: user_util
 * @Version: 1.0.0
 * @Date: 2022/1/8 20:23
 */

package user

import "github.com/ZimoBoy/gohub/pkg/database"

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断 Phone 已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
