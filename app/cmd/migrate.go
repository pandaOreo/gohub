/**
 * @Author: fanjinghua
 * @Description:
 * @File: migrate
 * @Version: 1.0.0
 * @Date: 2022/1/14 08:45
 */

package cmd

import (
	"github.com/ZimoBoy/gohub/migrations"
	"github.com/ZimoBoy/gohub/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run unmigrated migrations",
	// 所有 migrate 下的子命令都会执行以下代码
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
