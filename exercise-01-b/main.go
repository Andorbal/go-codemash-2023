package main

import (
	"fmt"
	"math"
)

func getValue(input int) int {
	if math.Sin(float64(input)) > 0 {
		return input
	}

	return 0
}

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += getValue(i)
	}

	fmt.Printf("Sum is %v\n", sum)
}
