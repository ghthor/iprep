package main

import (
	"fmt"
	"testing"
)

// InsertValues will insert the values in a reversed order
// to preserve the order presented in the slice. So the input
// [1,2,3,4,5] will be inserted such that {1 => 2 => 3 => 4 => 5 =>}
// the node for 1 contains a pointer to the node that contains 2,
// and so on.
func (n *IntList) InsertValues(values []int) *IntList {
	if len(values) == 0 {
		return n
	}

	if n == nil {
		n = &IntList{values[len(values)-1], nil}
		values = values[:len(values)-1]
	}

	for i := len(values) - 1; i >= 0; i-- {
		n = &IntList{values[i], n}
	}

	return n
}

// ContainsExactly will only return true if the list of
// nodes contains the same number of values in exact order
// as the input.
func (n *IntList) ContatinsExactly(values []int) (matches bool) {
	if len(values) == 0 && n == nil {
		return true
	}

	matches = true
	for _, v := range values {
		if n == nil {
			return false
		}

		matches = (n.v == v)

		if !matches {
			break
		}

		n = n.next
	}

	return
}

// Values will walk the list till the end and collect
// the values at each node into an array.
func (n *IntList) Values() []int {
	values := make([]int, 0, 10)

	for n != nil {
		values = append(values, n.v)
		n = n.next
	}

	return values
}

func (n *IntList) String() string {
	return fmt.Sprint(n.Values())
}

func TestIntListUnique(t *testing.T) {
	type testCase struct {
		input          []int
		expectedValues []int
	}

	cases := []testCase{{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	}, {
		[]int{1, 2, 1, 3, 1, 4, 4, 5},
		[]int{1, 2, 3, 4, 5},
	}}

	for i, c := range cases {
		var list *IntList = nil
		list = list.InsertValues(c.input)
		list = list.Unique()

		if !list.ContatinsExactly(c.expectedValues) {
			t.Logf("expected %v, got %v for case #%v", c.expectedValues, list, i)
			t.Fail()
		}
	}
}
