package functions

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumber := []int{5, 10, 15}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	tFn1 := getTransformerFunction(&numbers)
	tFn2 := getTransformerFunction(&moreNumber)

	num1 := transformNumbers(&numbers, tFn1)
	num2 := transformNumbers(&moreNumber, tFn2)

	fmt.Println(num1)
	fmt.Println(num2)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
