package lists

import "fmt"

func main() {
	prices := []float64{10.77, 22.43}

	prices = append(prices, 33.33)
	fmt.Println(prices)

	discountedPrices := []float64{5.5, 6.6, 7.7}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}

// func main() {
// 	productNames := [5]string{"Laptop", "Mouse", "Keyboard", "Monitor", "CPU"}
// 	prices := [4]float64{100.30, 200.10, 300.99, 400.60}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[0])

// 	// Slicing
// 	featuredPrices := prices[1:]
// 	featuredPrices[1] = 330.50
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
