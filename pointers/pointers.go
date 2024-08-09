package main

import "fmt"

func main() {
	age := 24

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("How many years I have been an adult: ",age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
