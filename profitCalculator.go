package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	const tax_rate float64 = 30

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	profit_before_tax := revenue - expenses
	fmt.Println("profit_before_tax: ", profit_before_tax)

	profit_after_tax := profit_before_tax * (1 - (tax_rate / 100))
	fmt.Println("profit_after_tax: ", profit_after_tax)

	fmt.Println("ratio: ", profit_before_tax/profit_after_tax)
}
