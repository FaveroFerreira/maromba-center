package db

import (
	"github.com/faveroferreira/maromba-center/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Trainer{})

	return db
}
