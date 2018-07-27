package exchange_rates_process

import (
	"github.com/JILeXanDR/parser/parsers"
	"github.com/JILeXanDR/parser/models"
	"math"
	"time"
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

// обрабатываем массив курсов покупки валюты для каждого сайта и возвращаем статистику
func CalculateStatisticData(allResults [][]parsers.CurrencyRateResult) ([]models.RateStatistic, error) {

	var (
		res = make([]models.RateStatistic, 0)
		min models.Money
		avg models.Money
		max models.Money
	)

	// перебираем каждую валюту, для каждой секунды графика
	for _, ratesPerCurrencyAndTime := range groupByTimeCurrency(allResults) {

		min = models.Money(math.Inf(1))
		avg = 0
		max = models.Money(math.Inf(-1))

		// TODO рассчитать среднюю, минимальну. максимальную цены покупки для каждого периода каждой валюты (1 сек)
		for _, rate := range ratesPerCurrencyAndTime.rates {
			if rate.Buy < min {
				min = rate.Buy
			}
			if rate.Buy > max {
				max = rate.Buy
			}
			avg += rate.Buy
		}

		payload := &models.RateStatistic{
			BaseCurrency:   models.BASE_CURRENCY,
			ToCurrency:     ratesPerCurrencyAndTime.currency,
			Time:           time.Unix(int64(ratesPerCurrencyAndTime.timestamp), 0),
			MinRate:        min.Round(),
			AvgRate:        models.Money(avg / models.Money(len(ratesPerCurrencyAndTime.rates))).Round(),
			MaxRate:        max.Round(),
			ResourcesCount: len(ratesPerCurrencyAndTime.rates),
		}

		res = append(res, *payload)
	}

	return res, nil
}
