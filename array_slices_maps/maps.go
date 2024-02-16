package main

import "fmt"

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"google": "https://google.com",
	// 	"amazon": "https://aws.com",
	// }
	// fmt.Println(websites)
	// fmt.Println(websites["amazon"])

	// websites["linkedIn"] = "https://linkedin.com"
	// fmt.Println(websites)

	// 2 = length, 5 = capacity
	userNames := make([]string, 2, 5)
	userNames[0] = "just"

	userNames = append(userNames, "subs")
	userNames = append(userNames, "admin")
	fmt.Println(userNames)

	// cannot set empty slots
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["rust"] = 3.4
	courseRatings["js"] = 4.5

	courseRatings.output()

	for _, rating := range courseRatings {
		fmt.Println(rating)
	}
}
