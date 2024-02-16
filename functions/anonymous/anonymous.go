package anonymous

import "fmt"

// anonymous function
func main() {
	numbers := []int{1, 2, 3}

	transformed := tranformNumerbers(&numbers, func(number int) int { return number * 2 })

	fmt.Println(transformed)

	doubleFn := createTransformer(2)
	number := doubleFn(10)
	fmt.Println(number)
}

func tranformNumerbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
