/**
 * @Author: fanjinghua
 * @Description:
 * @File: cache
 * @Version: 1.0.0
 * @Date: 2022/1/16 12:58
 */

package bootstrap

import (
	"fmt"
	"github.com/ZimoBoy/gohub/pkg/cache"
	"github.com/ZimoBoy/gohub/pkg/config"
)

// SetupCache 缓存
func SetupCache() {
	// 初始化缓存专用的 redis client, 使用专属缓存DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)
}
