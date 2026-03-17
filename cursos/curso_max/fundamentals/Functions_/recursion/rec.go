package main

import "fmt"

func fatorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Print(fatorial(5))
}