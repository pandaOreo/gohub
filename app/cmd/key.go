/**
 * @Author: fanjinghua
 * @Description:
 * @File: key
 * @Version: 1.0.0
 * @Date: 2022/1/13 10:34
 */

package cmd

import (
	"github.com/ZimoBoy/gohub/pkg/console"
	"github.com/ZimoBoy/gohub/pkg/helpers"
	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
