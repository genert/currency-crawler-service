package currency

import (
	"log"
	"time"
)

// ExchangeRateList represents a map key value of currency and rate.
type ExchangeRateList map[string]float64

type ExchangeRates struct {
	Timestamp *time.Time
	Rates     ExchangeRateList
}

type ExchangeCrawler interface {
	CrawlLatest() (ExchangeRates, error)
	CrawlByDate(time.Time) (ExchangeRates, error)
}

func UpdateCurrencyRates(crawler ExchangeCrawler) error {
	log.Print("starting currency rate update...")

	rates, err := crawler.CrawlLatest()
	if err != nil {
		return err
	}

	if rates.Timestamp == nil {
		log.Print("failed to get any rates")
		return nil
	}

	log.Print("retreive rates", len(rates.Rates))

	return nil
}
