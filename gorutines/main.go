package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go greet("Nice to meet you", done)
	go greet("How are you?", done)
	go slowGreet("How... are... you... ?", done)
	go greet("I hope you're doing fine?!", done)

	for range done {
		// fmt.Println(doneChan)
	}
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello! ", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hellow slow, ", phrase)
	doneChan <- true
	// has to be sure this operation takes the longest
	close(doneChan)
}
