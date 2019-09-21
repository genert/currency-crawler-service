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

type CurrencyRateRecord struct {
	ID            int       `json:"id"`
	Currency      string    `json:"currency"`
	Rate          float64   `json:"rate"`
	RateTimestamp time.Time `json:"rate_timestamp"`
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

	log.Print("retreived rates", len(rates.Rates))

	var records []CurrencyRateRecord
	for currency, rate := range rates.Rates {
		records = append(records, CurrencyRateRecord{
			Currency:      currency,
			Rate:          rate,
			RateTimestamp: *rates.Timestamp,
		})
	}

	// Here you can do whatever you want with
	// these rates such as save to database or whatever.
	log.Print(records)

	return nil
}
