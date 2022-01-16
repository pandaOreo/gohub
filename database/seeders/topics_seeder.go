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

	seed.Add("SeedTopicsTable", func(db *gorm.DB) {

		topics := factories.MakeTopics(10)

		result := db.Table("topics").Create(&topics)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
