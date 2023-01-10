package main

import "fmt"

func getOutput(input int) string {
	for i := 1; i < 20; i++ {
		if i%15 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(getOutput(i))
	}
}
