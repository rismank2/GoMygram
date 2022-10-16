package database

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "273161"
	dbPort   = "5432"
	dbname   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Gagal", err)
	} else {
		log.Println("Koneksi Berhasil")
	}
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Commment{}, models.SocialMedia{})
}
