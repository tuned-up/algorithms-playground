package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var max int

	flag.IntVar(&max, "max", 126, "Max number")
	flag.Parse()

	factors := make([]int, 0)

	for max % 2 == 0 {
		factors = append(factors, 2)
		max = max / 2
	}

	maxFactor := int(math.Ceil(math.Sqrt(float64(max))))
	for i := 3; i <= maxFactor; i = i + 2 {
		for max % i == 0 {
			factors = append(factors, i)
			max = max / i
			maxFactor = int(math.Ceil(math.Sqrt(float64(max))))
		}
	}

	if max > 1 {
		factors = append(factors, max)
	}

	fmt.Println("Factors are: ", factors)
}
