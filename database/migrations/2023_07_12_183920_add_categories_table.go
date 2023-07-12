package migrations

import (
	"database/sql"
	"github.com/lihc520/gohub/app/models"
	"github.com/lihc520/gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {
	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2023_07_12_183920_add_categories_table", up, down)
}
