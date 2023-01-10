package main

import "fmt"

func getValue(input int) int {
	if input%3 == 0 || input%5 == 0 {
		return input
	}

	return 0
}

func main() {
	sum := 0

	for i := 0; i < 1000; i++ {
		sum += getValue(i)
	}

	fmt.Printf("Sum is %v\n", sum)
}
