# Currency crawler service 

A currency crawler service written in Golang for AWS Lambda.

Currently, it fetches rates from ECB but You can add any other crawler you wish. Just implement a struct that satisfies ExchangeCrawler interface.

It does not do anything with the data other than print it. It is up to You to work on that - save to database or whatever is needed.

Enjoy!
