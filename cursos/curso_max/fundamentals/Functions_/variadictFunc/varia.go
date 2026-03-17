package main

import "fmt"

func somaLista(numbers []int) int {
	soma := 0
	for _, v := range numbers {
		soma += v
	}
	return soma
}

func main() {
	numbers := []int{2,2,3}
	somaArray := somaLista(numbers)
	fmt.Print(somaArray)
}