package main

import "fmt"

func main() {
	productNames := [5]string{"Laptop", "Mouse", "Keyboard", "Monitor", "CPU"}
	prices := [4]float64{100.30, 200.10, 300.99, 400.60}
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[0])

	// Slicing
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
