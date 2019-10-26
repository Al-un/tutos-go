package concurrency

import (
	"fmt"
	"time"
)

// a very simple function that we'll
// make asynchronous later on
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

// RunConcurrency https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
func RunConcurrency() {
	fmt.Println("Goroutine Tutorial")

	go func() {
		fmt.Println("Anonymous goroutines")
	}()

	// sequential execution of our compute function
	go compute(10)
	go compute(10)

	// we scan fmt for input and print that to our console
	var input string
	fmt.Scanln(&input)
}
