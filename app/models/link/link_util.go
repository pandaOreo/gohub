package link

import (
	"github.com/ZimoBoy/gohub/pkg/app"
	"github.com/ZimoBoy/gohub/pkg/cache"
	"github.com/ZimoBoy/gohub/pkg/database"
	"github.com/ZimoBoy/gohub/pkg/helpers"
	"github.com/ZimoBoy/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
	"time"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	// 取数据
	cache.GetObject(cacheKey, &links)

	// 如果数据为空
	if helpers.Empty(links) {
		// 查询数据库
		links = All()
		if helpers.Empty(links) {
			return links
		}
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
