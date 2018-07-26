package helpers

import (
	"github.com/JILeXanDR/parser/models"
	"strconv"
	"math/rand"
)

func StringToMoney(val string) models.Money {
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}

	return models.Money(i)
}

func RandFloat(min int, max int) float64 {
	return float64(min) + rand.Float64()*(float64(min)-float64(max))
}
