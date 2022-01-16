//Package toopic 模型
package toopic

import (
	"github.com/ZimoBoy/gohub/app/models"
	"github.com/ZimoBoy/gohub/app/models/category"
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/pkg/database"
)

type Toopic struct {
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

func (toopic *Toopic) Create() {
	database.DB.Create(&toopic)
}

func (toopic *Toopic) Save() (rowsAffected int64) {
	result := database.DB.Save(&toopic)
	return result.RowsAffected
}

func (toopic *Toopic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&toopic)
	return result.RowsAffected
}
