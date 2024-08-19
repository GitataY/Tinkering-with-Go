package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 2, 3, 4, 5)
	fmt.Println(sum)

}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val // sum = sum + val
	}
	return sum
}
