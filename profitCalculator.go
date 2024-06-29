package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//var revenue float64
	//var expenses float64
	const tax_rate float64 = 30

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		//panic("Invalid Revenue")
		return
	}

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	expenses, err := getUserInput("Enter expenses: ")

	if err != nil {
		fmt.Println(err)
		//panic("Invalid Expense")
		return
	}
	// profit_before_tax := revenue - expenses
	// fmt.Println("profit_before_tax: ", profit_before_tax)

	// profit_after_tax := profit_before_tax * (1 - (tax_rate / 100))
	// fmt.Println("profit_after_tax: ", profit_after_tax)

	// fmt.Println("ratio: ", profit_before_tax/profit_after_tax)
	pbt, pat, ratio := calculate(revenue, expenses, tax_rate)

	fmt.Println("profit_before_tax: ", pbt)
	fmt.Println("profit_after_tax: ", pat)
	fmt.Println("ratio: ", ratio)
	storeResults(pbt, pat, ratio)
}

func getUserInput(Text string) (float64, error) {
	var userInput float64
	fmt.Println(Text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid Input ,Please enter amount which is greater then zero")
	}
	return userInput, nil
}

func calculate(revenue float64, expenses float64, tax_rate float64) (float64, float64, float64) {
	profit_before_tax := revenue - expenses
	profit_after_tax := profit_before_tax * (1 - (tax_rate / 100))
	ratio := profit_before_tax / profit_after_tax
	return profit_before_tax, profit_after_tax, ratio
}

func storeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	err := os.WriteFile("results.txt", []byte(results), 0644)
	if err != nil {
		fmt.Println("Failed to write in file")
	}
}
