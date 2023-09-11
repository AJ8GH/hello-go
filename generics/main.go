package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}

func SumInts(ints map[string]int64) int64 {
	var sum int64
	for _, value := range ints {
		sum += value
	}
	return sum
}

func SumFloats(floats map[string]float64) float64 {
	var sum float64
	for _, value := range floats {
		sum += value
	}
	return sum
}

func SumIntsOrFloats[K comparable, V Number](numbers map[K]V) V {
	var sum V
	for _, value := range numbers {
		sum += value
	}
	return sum
}
