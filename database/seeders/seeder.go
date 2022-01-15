/**
 * @Author: fanjinghua
 * @Description:
 * @File: seeder
 * @Version: 1.0.0
 * @Date: 2022/1/15 13:23
 */

package seeders

import "github.com/ZimoBoy/gohub/pkg/seed"

func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法
	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUserTable",
	})
}
