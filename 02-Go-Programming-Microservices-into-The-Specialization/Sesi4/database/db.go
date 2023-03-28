package database

import (
	"fmt"
	"log"

	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi4/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "book"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(book.Book{})
}

func GetDB() *gorm.DB {
	return db
}