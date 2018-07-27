package exchange_rates_process

import (
	"github.com/JILeXanDR/parser/models"
	"github.com/JILeXanDR/parser/parsers"
	"log"
	"github.com/JILeXanDR/parser/db"
	"github.com/JILeXanDR/parser/ws"
)

// обработка результата парсинга для всех сайтов
func ProcessExchangeRates() {
	for {
		select {
		// результаты парсинга по всех источниках
		case allResults := <-parsers.ParserResultsChannel:

			payloads, _ := CalculateStatisticData(allResults)

			for _, payload := range payloads {
				err := db.Conn.Create(&payload).Error
				if err != nil {
					log.Println("Could not create RateStatistic record, ", err)
					return
				}

				// отправляем данные для обновления графика
				// каждое сообщения это отдельный отрезок времени
				if payload.ToCurrency == models.CURRENCY_USD {
					ws.SendMessageToAllClients(ws.NewWebSocketMessage("rates", []models.RateStatistic{payload}))
				}
			}
		}
	}
}
