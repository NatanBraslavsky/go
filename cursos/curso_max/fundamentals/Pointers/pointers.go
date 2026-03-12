package main

import "fmt"

func main() {
	age := 20
	getAge(&age)
	fmt.Print(age)
}

func getAge(age *int) {
	*age -= 10
}