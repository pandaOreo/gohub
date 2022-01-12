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
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户,通过 User.ID 来判断是否创建成功
func (u *User) Create() {
	database.DB.Create(&u)
}
