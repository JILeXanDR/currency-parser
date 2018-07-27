package parsers

import (
	"log"
	"time"
	"sync"
)

const DELAY_BEETWEEN_PARSING = 1 * time.Second

func GetSources() []string {
	return []string{
		"https://kurs.com.ua/gorod/2324-cherkassy/?utm_source=geo&utm_content=cherkassy",
		"https://minfin.com.ua/currency/banks/cherkassy/",
		"fake1",
		"fake2",
		"fake3",
		"fromFile",
	}
}

// изменяем время всех рузультатов
func changeTime(data [][]CurrencyRateResult) {
	for i1, v := range data {
		for i2, _ := range v {
			val := &data[i1][i2]
			val.Time = data[0][0].Time // FIXME panic: runtime error: index out of range (когда сласс пустой....)
		}
	}
}

func ParseAllSources(sourcesList []string) [][]CurrencyRateResult {

	var wg sync.WaitGroup
	var count = len(sourcesList)
	var allResults = make([][]CurrencyRateResult, 0, count)

	wg.Add(count)
	for _, url := range sourcesList {
		go func(url string) {
			defer wg.Done()
			results, err := parseUrl(url)
			if err != nil {
				log.Println(err)
				return
			}
			allResults = append(allResults, results)
		}(url)
	}
	// ждем пока завершится парсинг всех сайтов и запускаем парсер еще через 1 сек.
	wg.Wait()

	changeTime(allResults)

	return allResults
}

func RunParser() {
	for {
		ParserResultsChannel <- ParseAllSources(GetSources())
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
