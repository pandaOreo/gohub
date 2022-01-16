package main

/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 08:59:14
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 15:44:12
 * @Description  : 佛祖保佑,永无BUG
 */

import (
	"fmt"
	"github.com/ZimoBoy/gohub/app/cmd"
	"github.com/ZimoBoy/gohub/app/cmd/make"
	"github.com/ZimoBoy/gohub/bootstrap"
	btsConfig "github.com/ZimoBoy/gohub/config"
	"github.com/ZimoBoy/gohub/pkg/config"
	"github.com/ZimoBoy/gohub/pkg/console"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 应用的主入口, 默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数, --env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
