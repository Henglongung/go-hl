package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbName := os.Getenv("APP_DB_NAME")
	dbUser := os.Getenv("APP_DB_USER")
	dbPass := os.Getenv("APP_DB_PWD")
	dbProtocol := os.Getenv("APP_DB_PROTOCOL")

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbProtocol, dbName)
	log.Printf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbProtocol, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = database

}
