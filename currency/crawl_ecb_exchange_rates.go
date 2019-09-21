package currency

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

const EcbUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

type EcbEnvelope struct {
	Cube []EcbCube `xml:"Cube>Cube"`
}

type EcbCube struct {
	Date  string `xml:"time,attr"`
	Rates []struct {
		Currency string  `xml:"currency,attr"`
		Rate     float64 `xml:"rate,attr"`
	} `xml:"Cube"`
}

type crawler struct{}

func NewEcbExchangeRatesCrawler() ExchangeCrawler {
	return &crawler{}
}

func (c *crawler) CrawlLatest() (ExchangeRates, error) {
	return c.CrawlByDate(time.Now().AddDate(0, 0, -1).UTC())
}

// CrawlByDate returns currency list by date.
func (c *crawler) CrawlByDate(date time.Time) (ExchangeRates, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.UTC)

	resp, err := http.Get(EcbUrl)
	if err != nil {
		return ExchangeRates{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ExchangeRates{}, err
	}

	// Unmarshal into XML data.
	var exchangeResult EcbEnvelope
	err = xml.Unmarshal(data, &exchangeResult)
	if err != nil {
		return ExchangeRates{}, err
	}

	// Lets loop through result and save values to map.
	var rates = ExchangeRateList{}

	for _, c := range exchangeResult.Cube {
		// Check that date is in correct format like 2019-09-09 YYY-MM-DD
		if c.Date != date.Format("2006-01-02") {
			continue
		}

		for _, v := range c.Rates {
			rates[v.Currency] = v.Rate
		}
	}

	return ExchangeRates{
		Timestamp: &date,
		Rates:     rates,
	}, nil
}
