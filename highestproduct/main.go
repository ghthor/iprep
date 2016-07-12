package main

import "sort"

type Ints []int

type ByValueAscending []int

func (b ByValueAscending) Len() int {
	return len(b)
}

func (b ByValueAscending) Less(i int, j int) bool {
	return b[j] < b[i]
}

func (b ByValueAscending) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func (a Ints) HighestProduct() int {
	if len(a) < 3 {
		panic("list must contain at least 3 integer values")
	}

	// Sort the whole array
	sort.Sort(ByValueAscending(a))

	// Return the product of the 3 largest values
	return a[0] * a[1] * a[2]
}

func main() {
}
