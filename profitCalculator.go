package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	const tax_rate float64 = 30

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	revenue = getUserInput("Enter revenue: ")

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
	expenses = getUserInput("Enter expenses: ")

	// profit_before_tax := revenue - expenses
	// fmt.Println("profit_before_tax: ", profit_before_tax)

	// profit_after_tax := profit_before_tax * (1 - (tax_rate / 100))
	// fmt.Println("profit_after_tax: ", profit_after_tax)

	// fmt.Println("ratio: ", profit_before_tax/profit_after_tax)
	pbt, pat, ratio := calculate(revenue, expenses, tax_rate)
	fmt.Println("profit_before_tax: ", pbt)
	fmt.Println("profit_after_tax: ", pat)
	fmt.Println("ratio: ", ratio)
}

func getUserInput(Text string) float64 {
	var userInput float64
	fmt.Println(Text)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue float64, expenses float64, tax_rate float64) (float64, float64, float64) {
	profit_before_tax := revenue - expenses
	profit_after_tax := profit_before_tax * (1 - (tax_rate / 100))
	ratio := profit_before_tax / profit_after_tax
	return profit_before_tax, profit_after_tax, ratio
}
