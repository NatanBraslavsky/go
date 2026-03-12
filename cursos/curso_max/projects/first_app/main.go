package main

import "fmt"



func main() {
	revenue, expenses, taxRate := userInput()

	profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func userInput() (float64, float64, float64) {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter total revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Enter total expenses: ")
	fmt.Scanln(&expenses)

	fmt.Print("Enter tax rate (as a decimal): ")
	fmt.Scanln(&taxRate)

	return revenue, expenses, taxRate
} 

func calculate(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return profit, ratio
}