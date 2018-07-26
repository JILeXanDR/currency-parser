package parsers

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
	"github.com/JILeXanDR/parser/models"
	"regexp"
	"errors"
	"strconv"
)

type MinfinComUaParser struct {
	Url string
}

func (p *MinfinComUaParser) ParseUrl() []CurrencyRateResult {

	res, err := http.Get(p.Url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	parseRate := func(str string) (models.Money, error) {
		log.Println(str)
		var re = regexp.MustCompile(`(?m)([^\s]+)`)
		var res = re.FindStringIndex(str)

		if len(res) <= 0 {
			return 0, errors.New("Can not find rate")
		}

		strMoney := re.FindString(str)

		val, _ := strconv.ParseFloat(strMoney, 64)

		return models.Money(val), nil
	}

	var results = make([]CurrencyRateResult, 0)

	rows := doc.Find(".mfm-table tbody").First()
	rows.Each(func(i int, selection *goquery.Selection) {
		currencyEl := selection.Find(".mfcur-table-cur a").First()
		currencyKey := strings.TrimSpace(currencyEl.Text())

		if currencyKey == "ДОЛЛАР" {
			buyRate := selection.Find("td.mfm-text-nowrap").First().Text()
			rate, err := parseRate(buyRate)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, CurrencyRateResult{
				Currency: models.CURRENCY_USD,
				Buy:      rate,
				Time:     time.Now(),
			})
		}
	})

	return results
}
