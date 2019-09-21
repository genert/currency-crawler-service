package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"go.cantor.systems/currency-crawler-service/currency"
)

type Event struct {
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
	ecbExchangeRatesCrawler := currency.NewEcbExchangeRatesCrawler()
	err := currency.UpdateCurrencyRates(ecbExchangeRatesCrawler)
	if err != nil {
		return "nope", err
	}

	return "ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}
