/**
 * @Author: fanjinghua
 * @Description:
 * @File: app
 * @Version: 1.0.0
 * @Date: 2022/1/9 15:11
 */

package app

import "github.com/ZimoBoy/gohub/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
