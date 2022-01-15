/**
 * @Author: fanjinghua
 * @Description:
 * @File: app
 * @Version: 1.0.0
 * @Date: 2022/1/9 15:11
 */

package app

import (
	"github.com/ZimoBoy/gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间,支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带 v1 标识 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
