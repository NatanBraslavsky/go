package main

import "fmt"

func pares(numbers []int) []int {
	numPares := []int{}
	for _, v := range numbers {
		if v % 2 == 0 {
		numPares = append(numPares, v)
		}
	}
	return numPares
}

func main() {
	numbers := []int{2,4,5,55,67,32}
	numPares := pares(numbers)
	fmt.Println(numPares)
}

