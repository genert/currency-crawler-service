.PHONY: deps clean build run

deps:
	go get -u ./...

clean: 
	rm -rf ./main
	
build:
	GOOS=linux GOARCH=amd64 go build -o main ./

run:
	sam local invoke "CurrencyCrawlerService" -e currency-event.json
