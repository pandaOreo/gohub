/**
 * @Author: fanjinghua
 * @Description:
 * @File: make_factory
 * @Version: 1.0.0
 * @Date: 2022/1/15 22:37
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, exmaple: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1),
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	// 格式化模型名称,返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	// 基于模板创建文件 (做好变量替换)
	createFileFromStub(filePath, "factory", model)
}
