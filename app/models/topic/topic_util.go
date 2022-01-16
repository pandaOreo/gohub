package topic

import (
	"github.com/ZimoBoy/gohub/pkg/app"
	"github.com/ZimoBoy/gohub/pkg/database"
	"github.com/ZimoBoy/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (t Topic) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&t)
	return
}

func GetBy(field, value string) (t Topic) {
	database.DB.Where("? = ?", field, value).First(&t)
	return
}

func All() (t []Topic) {
	database.DB.Find(&t)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (t []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&t,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
