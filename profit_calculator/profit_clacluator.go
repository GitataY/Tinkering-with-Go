package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("A profit Calculator")

	revenue, err1 := outputText("revenue: ")
	expenses, err2 := outputText("Expenses: ")
	tax_rate, err3 := outputText(("Tax Rate: "))
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	earningsBeforeTax, profit, ratio := moniesmonies(revenue, expenses, tax_rate)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func storeResults(earningsBeforeTax, profit, ratio float64) {
	results := fmt.Sprintf("earningsBeforeTax: %.1f\nprofit: %.1f\nratio: %.3f\n", earningsBeforeTax, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func outputText(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func moniesmonies(revenue, expenses, tax_rate float64) (earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - tax_rate/100)
	ratio = earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}
