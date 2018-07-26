package repositories

import (
	"github.com/JILeXanDR/parser/db"
	"github.com/JILeXanDR/parser/models"
)

func LastStatistics(currency string) ([]models.RateStatistic, error) {
	payloads := make([]models.RateStatistic, 0)
	err := db.Conn.Limit(5).Order("time DESC").Find(&payloads, &models.RateStatistic{ToCurrency: currency}).Error
	if err != nil {
		return nil, err
	}

	return payloads, nil
}
