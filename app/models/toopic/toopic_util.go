package toopic

import (
	"github.com/ZimoBoy/gohub/pkg/app"
	"github.com/ZimoBoy/gohub/pkg/database"
	"github.com/ZimoBoy/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func Get(idstr string) (toopic Toopic) {
	database.DB.Where("id", idstr).First(&toopic)
	return
}

func GetBy(field, value string) (toopic Toopic) {
	database.DB.Where("? = ?", field, value).First(&toopic)
	return
}

func All() (toopics []Toopic) {
	database.DB.Find(&toopics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Toopic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (toopics []Toopic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Toopic{}),
		&toopics,
		app.V1URL(database.TableName(&Toopic{})),
		perPage,
	)
	return
}
