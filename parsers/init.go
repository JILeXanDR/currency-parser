package parsers

import (
	"github.com/JILeXanDR/parser/models"
	"strings"
	"time"
)

var (
	parseCurrencies = map[string]bool{models.CURRENCY_USD: true, models.CURRENCY_EUR: false}
	// результаты парсинга сайт/валюты
	// для 2 валют на 2 сайтах будет двумерный массив с 4 элементами
	ParserResultsChannel = make(chan [][]CurrencyRateResult, 0)
)

type CurrencyRateResult struct {
	Currency string
	Buy      models.Money
	Sale     models.Money
	Time     time.Time
}

type Parser interface {
	ParseUrl() []CurrencyRateResult
}

func GetParserForUrl(url string) (Parser, error) {

	var p Parser

	switch true {
	case strings.Contains(url, "kurs.com.ua"):
		p = &KursComUaParser{Url: url}
		break
	case strings.Contains(url, "minfin.com.ua"):
		p = &MinfinComUaParser{Url: url}
		break
	case strings.Contains(url, "fake1"):
		p = &Fake1Parser{}
		break
	case strings.Contains(url, "fake2"):
		p = &Fake2Parser{}
		break
	case strings.Contains(url, "fake3"):
		p = &Fake3Parser{}
		break
	case strings.Contains(url, "fromFile"):
		p = &FileParser{}
		break
	default:
		return nil, ErrCanNotFindParser
	}

	return p, nil
}
