package main

import "fmt"

// Variadic Functions
func main() {

	numbers := []int{1, 2, 3, 4}
	sum := sumup(numbers)
	sumTwo := sumUp(1, 2, 3, 4, 5, 6, 7)
	sumThree := sumUp(numbers...)

	fmt.Println(sum)
	fmt.Println(sumTwo)
	fmt.Println(sumThree)
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumUp(numbers ...int) int {

	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
