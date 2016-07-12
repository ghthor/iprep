package main

import "testing"

type highestProductCase struct {
	list            Ints
	expectedProduct int
}

func TestHighestProductSearch(t *testing.T) {

	cases := []highestProductCase{{
		[]int{1, 1, 1},
		1,
	}, {
		[]int{9, 5, 5, 4, 9, 6, 3, 4, 1, 10, 1, 6},
		9 * 9 * 10,
	}}

	for i, c := range cases {
		if c.list.HighestProduct() != c.expectedProduct {
			t.Logf("expected %v, got %v for case #%v", c.expectedProduct, c.list.HighestProduct(), i)
			t.Fail()
		}
	}
}
