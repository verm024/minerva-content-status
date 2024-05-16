package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbTcp := os.Getenv("DB_TCP")

	dsn := fmt.Sprintf("%s:%s%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbTcp, dbName)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("Connecting to database error: %v", err)
		return nil, err
	}
	log.Println("Success connecting to database")
	db.AutoMigrate()
	return db, nil
}
