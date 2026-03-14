package main

import "fmt"

func Contains[T comparable](numbers []T, n T) bool {
	for _, v := range numbers {
		if v == n {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{1,2,3,4}
	fmt.Print(Contains(numbers, 5))
}