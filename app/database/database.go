package database

import (
	"fmt"

	"github.com/koybigino/learn-fiber/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=Bielem@*01 dbname=fiber_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection Error")
	}

	fmt.Println("Connection succed !")

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		panic("Error ro create the table")
	}

	return db
}
