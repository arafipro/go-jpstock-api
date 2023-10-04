package config

import (
	"fmt"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const username string = "root"
const password string = "root"
const address string = "127.0.0.1:3306"
const dbname string = "go-jpstock"

func Connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, address, dbname,
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
