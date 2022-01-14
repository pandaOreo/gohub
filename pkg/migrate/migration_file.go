/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-13 21:49:10
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-13 21:55:46
 * @Description  : 佛祖保佑,永无BUG
 */

package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc 定义 up 和 down 回调方法的类型
type MigrationFunc func(gorm.Migrator, *sql.DB)

// migrationFiles 所有的迁移文件数组
var migrationFiles []MigrationFile

// MigrationFile 代表着带个迁移文件
type MigrationFile struct {
	Up       MigrationFunc
	Down     MigrationFunc
	FileName string
}

// Add 新增一个迁移文件, 所有的迁移文件都需要调用此方法来注册
func Add(name string, up MigrationFunc, down MigrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mfile MigrationFile) IsNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
