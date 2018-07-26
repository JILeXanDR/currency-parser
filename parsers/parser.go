package parsers

import (
	"log"
	"time"
	"sync"
)

const DELAY_BEETWEEN_PARSING = 1 * time.Second

func RunParser() {

	var sourcesList = []string{
		"https://kurs.com.ua/gorod/2324-cherkassy/?utm_source=geo&utm_content=cherkassy",
		"https://minfin.com.ua/currency/banks/cherkassy/",
		//"fake1",
		//"fake2",
		//"fake3",
	}

	var wg sync.WaitGroup
	var sourcesCount = len(sourcesList)
	var allResults = make([][]CurrencyRateResult, 0, sourcesCount)

	for {
		allResults = [][]CurrencyRateResult{}
		wg.Add(sourcesCount)
		for _, url := range sourcesList {
			go func(url string) {
				defer wg.Done()
				results, err := parseUrl(url)
				log.Println(results)
				if err != nil {
					log.Println(err)
					return
				}
				allResults = append(allResults, results)
			}(url)
		}
		// ждем пока завершится парсинг всех сайтов и запускаем парсер еще через 1 сек.
		wg.Wait()
		log.Println("All parsers done, len(allResults)=", len(allResults))
		ParserResultsChannel <- allResults
		time.Sleep(DELAY_BEETWEEN_PARSING)
	}
}

func parseUrl(url string) ([]CurrencyRateResult, error) {

	log.Println("Parse " + url + "...")

	parser, err := GetParserForUrl(url)
	if err != nil {
		return nil, err
	}

	var results = parser.ParseUrl()

	// отправляем данные о всех валютах для обработки
	return results, nil
}
