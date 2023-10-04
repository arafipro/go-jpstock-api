package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "root"
const password string = "root"
const address string = "127.0.0.1:3306"
const dbname string = "go-jpstock"

func Connect() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, address, dbname,
	)
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db = d

	// DBの接続確認
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
