package main

import "fmt"

type floatmap map[string]float64

func (f floatmap) output() {
	fmt.Println(f)
}

func main() {
	userName := make([]string, 2, 5)

	userName[0] = "Julie"

	userName = append(userName, "John")
	userName = append(userName, "Doe")

	fmt.Println(userName)

	courseRatings := make(floatmap, 3)

	courseRatings["Go Fundamentals"] = 5
	courseRatings["Docker Deep Dive"] = 4.5

	courseRatings.output()
	// fmt.Println(courseRatings)

}
