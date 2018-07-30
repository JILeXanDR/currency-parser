package db

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/JILeXanDR/parser/models"
	"fmt"
)

var Conn *gorm.DB

var (
	host     = "localhost"
	port     = 54320
	db       = "postgres"
	user     = "postgres"
	password = "d4REn0LdCH4B"
)

func InitDb() {

	db, err := gorm.Open("postgres", fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, db))
	if err != nil {
		log.Panicf("failed to connect database becauase in reason %s", err.Error())
	}

	Conn = db

	//Conn.LogMode(true)

	modelStructs := []interface{}{models.RateStatistic{}}

	//Conn.DropTable(modelStructs...)
	Conn.AutoMigrate(modelStructs...)
}
