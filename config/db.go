package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func InitDB() {
	db, err := gorm.Open(
		mysql.Open(os.Getenv("DATABASE_CONNECTION")),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	Conn = db
}

func GetDB() *gorm.DB {
	return Conn
}
