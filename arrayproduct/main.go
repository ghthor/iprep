package main

type ProductInputArray []int
type Products []int

func (p ProductInputArray) Products() Products {
	if len(p) == 0 {
		return []int{}
	}

	results := make(Products, 0, len(p))

	product := 1
	for i := range p {
		switch i {
		case 0:
			break

		default:
			product = product * p[i-1]
		}

		results = append(results, product)
	}

	product = 1
	for i := len(p) - 1; i >= 0; i-- {
		switch i {
		case len(p) - 1:
			continue

		default:
			product = product * p[i+1]
		}

		results[i] = results[i] * product
	}

	return results
}

func main() {
}
