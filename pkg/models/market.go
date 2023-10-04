package models

import (
	"github.com/arafipro/go-jpstock-api/pkg/config"
	"gorm.io/gorm"
)

var db_m *gorm.DB

type Market struct {
	MarketId uint8  `json:"market_id" gorm:"primaryKey"`
	Market   string `json:"market"`
}

func init() {
	config.Connect()
	db_m = config.GetDB()
	db_m.AutoMigrate(&Market{})
}
