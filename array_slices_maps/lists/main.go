package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!", "Get in shape!")
	fmt.Println(courseGoals)

	subGoals := []string{"make money", "run faster"}
	fmt.Println(subGoals)

	courseGoals = append(courseGoals, subGoals...)
	fmt.Println(courseGoals)

	// 7)
	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			129.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// func main() {
// 	// this creates a slice
// 	prices := []float64{10.99, 12.23, 32.43, 42.1}
// 	fmt.Println(prices[1])

// 	prices = append(prices, 42)
// 	fmt.Println(prices)

// }

// func main() {
// 	var productNames = [4]string{"Test Name"}
// 	prices := [4]float64{12.2, 9.99, 32.12, 45.4}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])
// 	productNames[2] = "A empty book"
// 	fmt.Println(productNames)

// slice - a subset of an array
// start from index 1 (include) to index 3 (exclude)
// slices is a reference to the original array
// featurePrices := prices[1:3]
// fmt.Println(featurePrices)
// spotPrice := prices[:2]
// fmt.Println(spotPrice)
// highPrices := prices[1:]
// fmt.Println(highPrices)

// // slice of slice
// slicePrice := highPrices[:1]
// fmt.Println(slicePrice)

// }
