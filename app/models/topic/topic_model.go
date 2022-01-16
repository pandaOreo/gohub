//Package topic 模型
package topic

import (
	"github.com/ZimoBoy/gohub/app/models"
	"github.com/ZimoBoy/gohub/app/models/category"
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/pkg/database"
)

type Topic struct {
	models.BaseModel

	// Put fields in here
	Title      string `json:"time,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (t *Topic) Create() {
	database.DB.Create(&t)
}

func (t *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&t)
	return result.RowsAffected
}

func (t *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&t)
	return result.RowsAffected
}
