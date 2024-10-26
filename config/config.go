package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	dbURI := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}
