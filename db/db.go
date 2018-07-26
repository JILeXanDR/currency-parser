package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/JILeXanDR/parser/models"
)

var Conn *gorm.DB

func InitDb() {
	var conn = fmt.Sprintf("host=localhost port=5433 user=postgres dbname=exchange_rates password=postgres")

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Panicf("failed to connect database becauase in reason %s", err.Error())
	}

	Conn = db

	//Conn.LogMode(true)

	modelStructs := []interface{}{models.ExchangeRate{}, models.RateStatistic{}}

	Conn.DropTable(modelStructs...)
	Conn.AutoMigrate(modelStructs...)
}
