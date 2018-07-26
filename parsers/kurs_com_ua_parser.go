package parsers

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
	"github.com/JILeXanDR/parser/helpers"
)

type KursComUaParser struct {
	Url string
}

func (p *KursComUaParser) ParseUrl() []CurrencyRateResult {

	res, err := http.Get(p.Url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results = make([]CurrencyRateResult, 0)

	items := doc.Find("#elCashBoard tr")
	items.Each(func(i int, selection *goquery.Selection) {
		currencyEl := selection.Find("td").First()
		currencyKey := strings.TrimSpace(currencyEl.Text())
		if parseCurrencies[currencyKey] {
			buyRate := selection.Find("td").Eq(1).Find(".ipsKurs_rate").Text()
			saleRate := selection.Find("td").Eq(3).Find(".ipsKurs_rate").Text()
			results = append(results, CurrencyRateResult{
				Currency: currencyKey,
				Buy:      helpers.StringToMoney(buyRate),
				Sale:     helpers.StringToMoney(saleRate),
				Time:     time.Now(),
			})
		}
	})

	return results
}
