package parsers

import (
	"github.com/JILeXanDR/parser/helpers"
	"time"
	"github.com/JILeXanDR/parser/models"
)

type Fake1Parser struct {
}

type Fake2Parser struct {
}

type Fake3Parser struct {
}

// min=27
// avg=28
// max=29

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
