package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB(dbSource string) {
	var err error

	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
	}
	Db = db
}

func GetDB() *gorm.DB {
	return Db
}