package main

type ProductInputArray []int
type Products []int

type indexRange struct {
	start, end int
}

// The product of a contigous index range.
type contiguousProduct struct {
	indexRange

	inputs  []int
	product int
}

func newContigousRangeProduct(start, end int, values []int) contiguousProduct {
	result := contiguousProduct{
		indexRange: indexRange{start, end},
		inputs:     values,
	}

	product := 1
	for _, v := range values {
		product = product * v
	}

	result.product = product

	return result
}

func (result contiguousProduct) Join(other contiguousProduct) int {
	return result.product * other.product
}

func (p ProductInputArray) Products() Products {
	products := make(Products, 0, len(p))

	for i := range p {
		result := 0

		switch i {
		case 0:
			result = newContigousRangeProduct(1, len(p), p[1:]).product

		case len(p) - 1:
			result = newContigousRangeProduct(0, len(p)-1, p[:len(p)-1]).product

		default:
			result = newContigousRangeProduct(0, i, p[0:i]).Join(newContigousRangeProduct(i+1, len(p), p[i+1:len(p)]))
		}

		products = append(products, result)
	}

	return products
}

func main() {
}
