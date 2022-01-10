/**
 * @Author: fanjinghua
 * @Description:
 * @File: redis
 * @Version: 1.0.0
 * @Date: 2022/1/10 22:43
 */

package config

import "github.com/ZimoBoy/gohub/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 业务类存储使用 1 (图片验证码,短信验证码,会话)
			"database": config.Env("REDIS_MAIN_DB", 1),
		}
	})
}
