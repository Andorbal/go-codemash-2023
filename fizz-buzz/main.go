package main

import "fmt"

func getOutput(input int) string {
	if input%15 == 0 {
		return "fizz buzz"
	} else if input%3 == 0 {
		return "fizz"
	} else if input%5 == 0 {
		return "buzz"
	} else {
		return fmt.Sprint(input)
	}
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(getOutput(i))
	}
}
