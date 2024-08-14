package main

import "fmt"

type Product struct {
	title string
	id	int
	price float64
}

func main() {
	hobbies := [3]string{"reading", "singing", "cooking"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1], hobbies[2])

	featuredHobbies := hobbies[0:2]
	fmt.Println(featuredHobbies)
	feturedHobbies2 := hobbies[:2]
	fmt.Println(feturedHobbies2)

	task3 := featuredHobbies[1:3]
	fmt.Println(task3)

	coursegoals := []string{"learn go", "learn python"}
	fmt.Println(coursegoals)
	coursegoals[1] = "learn java"
	coursegoals = append(coursegoals, "learn javascript")
	fmt.Println(coursegoals)

	ProductItems := []Product{ {
		title: "book",
		id: 1,
		price: 100.00,
	},
	{
		title: "pen",
		id: 2,
		price: 10.00,
	},
	}  
	fmt.Println(ProductItems)

	ProductItems = append(ProductItems, Product{
		title: "pencil",
		id: 3,
		price: 5.00,
	},)
	fmt.Println(ProductItems)
}


