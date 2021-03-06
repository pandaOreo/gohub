/**
 * @Author: fanjinghua
 * @Description:
 * @File: user_model
 * @Version: 1.0.0
 * @Date: 2022/1/8 20:21
 */

package user

import (
	"github.com/ZimoBoy/gohub/app/models"
	"github.com/ZimoBoy/gohub/pkg/database"
	"github.com/ZimoBoy/gohub/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户,通过 User.ID 来判断是否创建成功
func (u *User) Create() {
	database.DB.Create(&u)
}

// ComparePassword 判断密码是否正确
func (u *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, u.Password)
}

// Save 保存
func (u *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&u)
	return result.RowsAffected
}
