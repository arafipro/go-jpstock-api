package models

import (
	"github.com/arafipro/go-jpstock-api/pkg/config"
	"gorm.io/gorm"
)

var db_s *gorm.DB

type Stock struct {
	Code      uint8  `json:"code" gorm:"primaryKey"`
	Stockname string `json:"stockname"`
	MarketId  uint8  `json:"market_id"`
}

func init() {
	config.Connect()
	db_s = config.GetDB()
	db_s.AutoMigrate(&Stock{})
}
