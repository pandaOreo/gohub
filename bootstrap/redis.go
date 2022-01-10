/**
 * @Author: fanjinghua
 * @Description:
 * @File: redis
 * @Version: 1.0.0
 * @Date: 2022/1/10 22:45
 */

package bootstrap

import (
	"fmt"
	"github.com/ZimoBoy/gohub/pkg/config"
	"github.com/ZimoBoy/gohub/pkg/redis"
)

func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
