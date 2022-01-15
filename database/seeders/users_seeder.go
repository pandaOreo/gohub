/**
 * @Author: fanjinghua
 * @Description:
 * @File: users_seeder
 * @Version: 1.0.0
 * @Date: 2022/1/15 13:20
 */

package seeders

import (
	"fmt"
	"github.com/ZimoBoy/gohub/database/factories"
	"github.com/ZimoBoy/gohub/pkg/console"
	"github.com/ZimoBoy/gohub/pkg/logger"
	"github.com/ZimoBoy/gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 添加 Seeder
	seed.Add("SeedUserTable", func(db *gorm.DB) {
		// 创建 10 个用户对象
		users := factories.MakeUsers(10)

		// 批量创建用户 (注意批量创建不会调用模型钩子)
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
