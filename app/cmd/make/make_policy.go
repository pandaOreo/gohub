/**
 * @Author: fanjinghua
 * @Description:
 * @File: make_policy
 * @Version: 1.0.0
 * @Date: 2022/1/16 10:24
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	os.Mkdir("app/policies", os.ModePerm)

	filePath := fmt.Sprintf("app/policies/%s_policy.go", model.PackageName)

	createFileFromStub(filePath, "policy", model)
}
