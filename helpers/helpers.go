package helpers

import (
	"github.com/JILeXanDR/parser/models"
	"strconv"
)

func StringToMoney(val string) models.Money {
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}

	return models.Money(i)
}
