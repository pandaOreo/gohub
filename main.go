/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-08 08:59:14
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-08 14:52:25
 * @Description  : 佛祖保佑,永无BUG
 */

package main

import (
	"fmt"

	"github.com/ZimoBoy/gohub/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 GIN Engine 实例
	r := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务
	err := r.Run(":8000")
	if err != nil {
		// 错误处理, 端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
