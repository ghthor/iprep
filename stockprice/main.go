package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Prices []int

func (p Prices) MaxProfit() (profit int) {
	// TODO: Implementation
	return 0
}

func NewPricesFromStrings(pricesStr []string) (Prices, error) {
	prices := make(Prices, 0, len(pricesStr))
	for _, s := range pricesStr {
		p, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		prices = append(prices, p)
	}

	return prices, nil
}

func main() {
	pricesStr := flag.String("prices", "10,7,5,8,11,9", "A comma seperated list of stock prices")

	flag.Parse()

	prices, err := NewPricesFromStrings(strings.Split(*pricesStr, ","))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prices.MaxProfit())
}
