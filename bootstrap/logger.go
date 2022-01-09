/**
 * @Author: fanjinghua
 * @Description:
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2022/1/9 19:12
 */

package bootstrap

import (
	"github.com/ZimoBoy/gohub/pkg/config"
	"github.com/ZimoBoy/gohub/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
