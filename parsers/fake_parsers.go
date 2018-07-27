package parsers

import (
	"github.com/JILeXanDR/parser/helpers"
	"time"
	"github.com/JILeXanDR/parser/models"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Fake1Parser struct {
}

type Fake2Parser struct {
}

type Fake3Parser struct {
}

type FileParser struct {
}

func (p *Fake1Parser) ParseUrl() []CurrencyRateResult {
	return []CurrencyRateResult{
		{
			Currency: models.CURRENCY_USD,
			Buy:      helpers.StringToMoney("27"),
			Time:     time.Now(),
		},
	}
}

func (p *Fake2Parser) ParseUrl() []CurrencyRateResult {
	return []CurrencyRateResult{
		{
			Currency: models.CURRENCY_USD,
			Buy:      helpers.StringToMoney("28"),
			Time:     time.Now(),
		},
	}

}

func (p *Fake3Parser) ParseUrl() []CurrencyRateResult {
	return []CurrencyRateResult{
		{
			Currency: models.CURRENCY_USD,
			Buy:      helpers.StringToMoney("29"),
			Time:     time.Now(),
		},
	}
}

func (p *FileParser) ParseUrl() []CurrencyRateResult {

	file, err := ioutil.ReadFile("./currency_rates.json")
	if err != nil {
		log.Println(err)
		return []CurrencyRateResult{}
	}

	var jsontype map[string]models.Money
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)

	return []CurrencyRateResult{
		{
			Currency: models.CURRENCY_USD,
			Buy:      jsontype["USD"],
			Time:     time.Now(),
		},
		{
			Currency: models.CURRENCY_EUR,
			Buy:      jsontype["UER"],
			Time:     time.Now(),
		},
	}
}
