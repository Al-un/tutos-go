package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Printf("Calculated Random Value: %d\n", value)
	time.Sleep(1000 * time.Millisecond)
	values <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

// RunChannel https://tutorialedge.net/golang/go-channels-tutorial/
func RunChannel() {

	fmt.Println("Go Channel Tutorial")

	values := make(chan int, 2)
	defer close(values)

	go calculateValue(values)
	go calculateValue(values)

	value := <-values
	fmt.Printf("Final values are %v\n", value)
}
