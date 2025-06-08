package migrations

import (
	"log"

	"github.com/cmartinez/GolangRestApiTesting/internal/domain"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&domain.User{}, &domain.Task{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully!")
}
