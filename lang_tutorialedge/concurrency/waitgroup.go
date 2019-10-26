package concurrency

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Inside a goroutine")
	wg.Done()
}

// RunWaitGroup https://tutorialedge.net/golang/go-waitgroup-tutorial/
//
// Real world example skipped
func RunWaitGroup() {
	fmt.Println("Tuto WaitGroup: Start!")

	// init
	var wg sync.WaitGroup

	// wait group for function
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()

	// wait group for anonymous function
	wg.Add(1)
	go func() {
		fmt.Println("Inside an anonymous gorouting")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Tuto WaitGroup: End.")
}
