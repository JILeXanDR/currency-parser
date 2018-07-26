package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
	"github.com/JILeXanDR/parser/db"
	"github.com/JILeXanDR/parser/server"
	"github.com/JILeXanDR/parser/exchange_rates_process"
	"github.com/JILeXanDR/parser/parsers"
)

const SERVER_PORT = 12345;

func main() {

	db.InitDb()
	defer db.Conn.Close()

	// запускаем парсер валют со всех сайтов
	go parsers.RunParser()

	// обработка данных персера валют
	go exchange_rates_process.ProcessExchangeRates()

	server.InitServer(":" + strconv.Itoa(SERVER_PORT))
}
