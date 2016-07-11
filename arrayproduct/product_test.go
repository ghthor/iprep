package main

import "testing"

type productTestCase struct {
	input    ProductInputArray
	expected Products
}

func (p Products) Equals(other Products) bool {
	if len(p) != len(other) {
		return false
	}

	for i := range other {
		if p[i] != other[i] {
			return false
		}
	}

	return true
}

func TestArrayProductCalculation(t *testing.T) {
	if !(Products{1, 2, 3}).Equals(Products{1, 2, 3}) {
		t.Fatal("expected {1,2,3} == {1,2,3}")
	}

	if (Products{1, 2, 3, 4}).Equals(Products{1, 2, 3}) {
		t.Fatal("expected {1,2,3,4} != {1,2,3}")
	}

	cases := []productTestCase{{
		[]int{1, 7, 3, 4},
		[]int{84, 12, 28, 21},
	}}

	for i, c := range cases {
		if c.input.Products().Equals(c.expected) {
			continue
		}

		t.Logf("expected %v, got %v for case #%v", c.expected, c.input.Products(), i)
		t.Fail()
	}
}
