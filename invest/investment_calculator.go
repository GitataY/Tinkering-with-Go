package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)



	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	futureRealValue := futureValue / math.Pow(1+ inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues() {

}