package main

import "testing"

type ProfitCase struct {
	prices         Prices
	expectedProfit int
}

func TestMaxProfitCalculation(t *testing.T) {
	cases := []ProfitCase{{
		[]int{10, 7, 5, 8, 11, 9},
		6,
	}, {
		[]int{5, 5, 5, 5, 5, 5, 5, 5},
		0,
	}, {
		[]int{10, 9, 8, 7, 3, 1},
		-1,
	}}

	for i, c := range cases {
		profit := c.prices.MaxProfit()
		if profit != c.expectedProfit {
			t.Logf("case %v: expected profit == %v, got == %v", i, c.expectedProfit, profit)
			t.Fail()
		}
	}
}
