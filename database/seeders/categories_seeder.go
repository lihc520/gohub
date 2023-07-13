package seeders

import (
    "fmt"
    "github.com/lihc520/gohub/database/factories"
    "github.com/lihc520/gohub/pkg/console"
    "github.com/lihc520/gohub/pkg/logger"
    "github.com/lihc520/gohub/pkg/seed"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedCategoriesTable", func(db *gorm.DB) {

        categories := factories.MakeCategories(10)

        result := db.Table("categories").Create(&categories)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}