/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-13 21:56:31
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-13 21:56:37
 * @Description  : 佛祖保佑,永无BUG
 */

package migrate

import (
	"github.com/ZimoBoy/gohub/pkg/database"
	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例, 用以执行迁移操作
func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在的话就创建他
	migrator.createMigrationsTable()
	return migrator
}

func (m *Migrator) createMigrationsTable() {
	migration := Migrator{}

	// 不存在才创建
	if !m.Migrator.HasTable(&migration) {
		m.Migrator.CreateTable(&migration)
	}
}
