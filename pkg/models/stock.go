package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arafipro/go-jpstock/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

var db_s *sql.DB

type Stock struct {
	Code      int8   `json:"code"`
	Stockname string `json:"stockname"`
	MarketId  int8   `json:"market_id"`
}

func CreateStockTable() {
	tableName := "stocks"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			code int(8) PRIMARY KEY,
			stockname varchar(128) NOT NULL,
			market_id int(8) NOT NULL)`, tableName)
	var err error
	_, err = db_s.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	config.Connect()
	db_s = config.GetDB()
}
