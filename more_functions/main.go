package main

import "fmt"

type transform func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// moreNumbers := []int{6, 7, 8, 9, 10}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform transform) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transform {
	// This function returns a function
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

func double(number int) int {
	return number * 2
}
func triple(number int) int {
	return number * 3
}
