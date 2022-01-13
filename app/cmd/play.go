/**
 * @Author: fanjinghua
 * @Description:
 * @File: play
 * @Version: 1.0.0
 * @Date: 2022/1/13 10:49
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// runPlay 调试完成后请记得清楚测试代码
func runPlay(cmd *cobra.Command, args []string) {
	//// 存进去 redis 中
	//redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	//// 从redis里取出
	//console.Success(redis.Redis.Get("hello"))
}
