/**
 * @Author: fanjinghua
 * @Description:
 * @File: make_seeder
 * @Version: 1.0.0
 * @Date: 2022/1/15 16:03
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example: make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeeder(cmd *cobra.Command, args []string) {
	// 格式化模型名称, 返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)

	// 基于模板创建文件 (做好变量替换)
	createFileFromStub(filePath, "seeder", model)
}
