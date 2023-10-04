package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arafipro/go-jpstock/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

var db_m *sql.DB

type Market struct {
	MarketId int8   `json:"market_id"`
	Market   string `json:"market"`
}

func CreateMarketTable() {
	tableName := "markets"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			market_id int(8) PRIMARY KEY,
			market varchar(128) NOT NULL)`, tableName)
	var err error
	_, err = db_m.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	config.Connect()
	db_m = config.GetDB()
}
