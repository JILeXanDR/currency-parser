package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
)

const (
	CURRENCY_USD = "USD"
	CURRENCY_UAH = "UAH"
	CURRENCY_EUR = "EUR"
)

const BASE_CURRENCY = CURRENCY_UAH

type Money float32

// округляем значение к сотым
func (m Money) Round() Money {
	return Money(math.Floor(float64(m)*100) / 100)
}

// создает тип с строки, например "15.05" будет как Money(15.05)
func NewMoneyFromString(val string) Money {
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}

	return Money(i)
}

// запись о курсе валют
//type ExchangeRate struct {
//	gorm.Model
//	BaseCurrency string    `json:"base_currency"`
//	ToCurrency   string    `json:"to_currency"`
//	Source       string    `json:"source"`
//	Buy          Money     `json:"buy"`
//	Sale         Money     `json:"sale"`
//	FoundAt      time.Time `json:"found_at"`
//}

type RateStatistic struct {
	gorm.Model
	BaseCurrency   string    `json:"base_currency"`
	ToCurrency     string    `json:"base_currency"`
	MinRate        Money     `json:"min_rate"`
	AvgRate        Money     `json:"avg_rate"`
	MaxRate        Money     `json:"max_rate"`
	ResourcesCount int       `json:"resources_count"`
	Time           time.Time `json:"time"`
}
