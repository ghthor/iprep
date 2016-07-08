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
	if len(p) < 2 {
		panic("len(prices array) < 2")
	}

	// Set initial price seen
	highestPriceSeen := p[len(p)-1]
	highestProfit := p[len(p)-1] - p[len(p)-2]

	for i := len(p) - 2; i >= 0; i-- {
		if p[i] > highestPriceSeen {
			remaining := p[:i+1]
			if len(remaining) < 2 {
				continue
			}

			profit := remaining.MaxProfit()
			if profit > highestProfit {
				return profit
			} else {
				highestPriceSeen = p[i]
				continue
			}

		} else {
			profit := highestPriceSeen - p[i]
			if profit > highestProfit {
				highestProfit = profit
			}
		}
	}

	return highestProfit
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
