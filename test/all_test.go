package test

import (
	"testing"
	"github.com/JILeXanDR/parser/parsers"
	"github.com/JILeXanDR/parser/exchange_rates_process"
	"github.com/stretchr/testify/assert"
	"github.com/JILeXanDR/parser/models"
	"time"
)

func createFakeRates(values ... models.Money) [][]parsers.CurrencyRateResult {

	resultsPerSource := make([][]parsers.CurrencyRateResult, 3)

	for _, value := range values {
		resultsPerSource = append(resultsPerSource, []parsers.CurrencyRateResult{
			{
				Currency: models.CURRENCY_USD,
				Buy:      value,
				Time:     time.Now(),
			},
		})
	}

	return resultsPerSource
}

// тестирование правильности подсчета минимальной. средней, максимальной цены покупки валюты с 5 сайтов
func TestCalculateStatisticDataForOneCurrencyFromDifferentSites(t *testing.T) {

	initialTime := time.Now()

	allResults := [][]parsers.CurrencyRateResult{
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.50), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.60), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.70), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(26.05), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(27.00), Time: initialTime},
		},
	}

	payloads, _ := exchange_rates_process.CalculateStatisticData(allResults)

	assert.Len(t, payloads, 1, "Для одной валюты в один промежуток времени (1 сек) должна быть одна запись статистики")

	// USD
	rate := payloads[0]

	assert.Equal(t, models.Money(25.5), rate.MinRate, "Минимальная цена неверная")
	assert.Equal(t, models.Money(25.97), rate.AvgRate, "Средняя цена неверная")
	assert.Equal(t, models.Money(27), rate.MaxRate, "Мксимальная цена неверная")
}

func TestCalculateStatisticDataForTwoCurrencyFromDifferentSites(t *testing.T) {
	initialTime := time.Now()

	allResults := [][]parsers.CurrencyRateResult{
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.50), Time: initialTime},
			{Currency: models.CURRENCY_EUR, Buy: models.Money(30), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.60), Time: initialTime},
			{Currency: models.CURRENCY_EUR, Buy: models.Money(30.20), Time: initialTime},
		},
		{
			{Currency: models.CURRENCY_USD, Buy: models.Money(25.70), Time: initialTime},
			{Currency: models.CURRENCY_EUR, Buy: models.Money(31), Time: initialTime},
		},
	}

	payloads, _ := exchange_rates_process.CalculateStatisticData(allResults)

	assert.Len(t, payloads, 2, "Для одной валюты в один промежуток времени (1 сек) должна быть одна запись статистики")

	rateUsd := payloads[0]
	rateEur := payloads[1]

	assert.Equal(t, models.Money(25.5), rateUsd.MinRate, "Минимальная цена неверная (USD)")
	assert.Equal(t, models.Money(25.6), rateUsd.AvgRate, "Средняя цена неверная (USD)")
	assert.Equal(t, models.Money(25.7), rateUsd.MaxRate, "Максимальная цена неверная (USD)")

	assert.Equal(t, models.Money(30), rateEur.MinRate, "Минимальная цена неверная (EUR)")
	assert.Equal(t, models.Money(30.4), rateEur.AvgRate, "Средняя цена неверная (EUR)")
	assert.Equal(t, models.Money(31), rateEur.MaxRate, "Максимальная цена неверная (EUR)")
}

func TestParserReturnsValidData(t *testing.T) {
	allResults := parsers.ParseAllSources([]string{"fake1", "fake2", "fake3"})
	assert.Equal(t, 3, len(allResults), "Парсинг 3 источников должен вернуть 3 значения")
	//for _, res := range allResults {
	//	log.Println(res)
	//}
}

func TestCreateMoneyFromString(t *testing.T) {
	assert.Equal(t, 13.33, float64(models.NewMoneyFromString("13.33")))
	assert.Equal(t, 0.50, float64(models.NewMoneyFromString("0.5")))
	assert.Equal(t, 2.5, float64(models.NewMoneyFromString("2.50")))
}

func TestRoundMoney(t *testing.T) {
	assert.Equal(t, 1.53, float64(models.Money(1.533).Round()))
}
