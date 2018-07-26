package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

const (
	CURRENCY_USD = "USD"
	CURRENCY_UAH = "UAH"
	CURRENCY_EUR = "EUR"
)

const BASE_CURRENCY = CURRENCY_UAH

type Money float32

// запись о курсе валют
type ExchangeRate struct {
	gorm.Model
	BaseCurrency string    `json:"base_currency"`
	ToCurrency   string    `json:"to_currency"`
	Source       string    `json:"source"`
	Buy          Money     `json:"buy"`
	Sale         Money     `json:"sale"`
	FoundAt      time.Time `json:"found_at"`
}

type RateStatistic struct {
	gorm.Model
	BaseCurrency string    `json:"base_currency"`
	ToCurrency   string    `json:"base_currency"`
	MinRate      Money     `json:"min_rate"`
	AvgRate      Money     `json:"avg_rate"`
	MaxRate      Money     `json:"max_rate"`
	Time         time.Time `json:"time"`
}
