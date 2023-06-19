package db

import (
	_ "database/sql"
	"log"
	"os"
	"project-inventory-book/models"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{
		{
			Title:       "Perjalanan Ini",
			Author:      "Andy",
			Description: "Buku tentang perjalanan",
			Stock:       10,
		},
		{
			Title:       "Pengobatan Ini",
			Author:      "Budy",
			Description: "Buku tentang pengobatan",
			Stock:       15,
		}, {
			Title:       "Hewan Ini",
			Author:      "Cici",
			Description: "Buku tentang hewan",
			Stock:       20,
		}, {
			Title:       "Laut Ini",
			Author:      "Didi",
			Description: "Buku tentang laut",
			Stock:       25,
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}
