package main

import (
	"fmt"
)

func getValue(input int) int {
	if input%2 == 0 {
		return input
	}

	return 0
}

func main() {
	sum := 0

	termLow := 1

	for termHigh := termLow + termLow; termHigh <= 4_000_000; {
		sum += getValue(termHigh)
		temp := termHigh
		termHigh += termLow
		termLow = temp
	}

	fmt.Printf("Sum is %v\n", sum)
}
