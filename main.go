package main

/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 08:59:14
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 15:44:12
 * @Description  : 佛祖保佑,永无BUG
 */

import (
	"flag"
	"fmt"

	"github.com/ZimoBoy/gohub/bootstrap"
	btsConfig "github.com/ZimoBoy/gohub/config"
	"github.com/ZimoBoy/gohub/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化, 依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件, 如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run("localhost:" + config.Get("app.port"))
	if err != nil {
		// 错误处理, 端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
