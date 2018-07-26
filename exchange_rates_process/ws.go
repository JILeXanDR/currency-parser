package exchange_rates_process

import (
	"github.com/JILeXanDR/parser/models"
	"time"
	"github.com/JILeXanDR/parser/parsers"
	"log"
	"math"
	"github.com/JILeXanDR/parser/db"
	"github.com/JILeXanDR/parser/ws"
)

// сгрупированная информация по всех источникам для каждой валюты в отдельный промежуток времени
type group struct {
	timestamp int
	currency  string
	rates     []parsers.CurrencyRateResult
}

// найти такое время и валюту и добавить в массив rates
func appendByCurrencyAndTimestamp(g []group, rate parsers.CurrencyRateResult) []group {
	var timestamp = int(rate.Time.Unix())

	for index, _ := range g {
		v := &g[index]
		if v.timestamp == timestamp && v.currency == rate.Currency {
			v.rates = append(v.rates, rate)
			return g
		}
	}

	// новая группа
	return append(g, group{
		currency:  rate.Currency,
		timestamp: timestamp,
		rates:     []parsers.CurrencyRateResult{rate},
	})
}

func groupByTimeCurrency(results [][]parsers.CurrencyRateResult) []group {

	var timeGroup = make([]group, 0)

	for _, perSite := range results {
		for _, rate := range perSite {
			timeGroup = appendByCurrencyAndTimestamp(timeGroup, rate)
		}
	}

	return timeGroup
}

// обработка курса валют каждые 3 сек
func ProcessExchangeRates() {

	//var err error

	for {
		select {
		// результаты парсинга по всех источниках
		case allResults := <-parsers.ParserResultsChannel:

			log.Println("Found allResults=", len(allResults))

			// перебираем каждую валюту, для каждой секунды графика
			for _, ratesPerCurrencyAndTime := range groupByTimeCurrency(allResults) {

				var min models.Money = math.MaxInt64
				var avg models.Money = 0
				var max models.Money = math.MinInt64

				if len(ratesPerCurrencyAndTime.rates) != 3 {
					//log.Panicf("Неправильное количество валют %v!", len(ratesPerCurrencyAndTime.rates))
				}

				for _, rate := range ratesPerCurrencyAndTime.rates {
					// TODO рассчитать среднюю, минимальну. максимальную цены покупки для каждого периода каждой валюты (1 сек)
					if rate.Buy < min {
						min = rate.Buy
					}
					if rate.Buy > max {
						max = rate.Buy
					}
					avg += rate.Buy
				}

				payload := &models.RateStatistic{
					BaseCurrency: models.BASE_CURRENCY,
					ToCurrency:   ratesPerCurrencyAndTime.currency,
					Time:         time.Unix(int64(ratesPerCurrencyAndTime.timestamp), 0),
					MinRate:      min,
					AvgRate:      avg / models.Money(len(ratesPerCurrencyAndTime.rates)),
					MaxRate:      max,
				}

				err := db.Conn.Create(payload).Error
				if err != nil {
					log.Println("Could not create RateStatistic record, ", err)
					return
				}

				// TODO remove
				if !(payload.MinRate == 27 && payload.AvgRate == 28 && payload.MaxRate == 29) {
					log.Printf("STATISTIC => %v %v %v", payload.MinRate, payload.AvgRate, payload.MaxRate)
				}

				// отправляем данные для обновления графика
				// каждое сообщения это отдельный отрезок времени
				if payload.ToCurrency == models.CURRENCY_USD {
					ws.SendMessageToAllClients(ws.NewWebSocketMessage("rates", []models.RateStatistic{*payload}))
				}
			}
		}
	}
}
