package main

import (
	"fmt"
	"sync"
	"time"
)

func writeNumber(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Writing %d to channel\n", i)
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Goroutine ", i)
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("Ended %d goroutine\n", i)
	wg.Done()
}

func main() {
	// --- Simple buffered channel
	ch := make(chan string, 2)
	ch <- "Bob"
	ch <- "Jack"
	// ch <- "Joe" // this line will trigger `fatal error: all goroutines are asleep - deadlock`
	fmt.Printf("Capacity %d\n", cap(ch))
	fmt.Printf("Length %d\n", len(ch))
	fmt.Printf("Content %s\n", <-ch)
	fmt.Printf("Length %d\n", len(ch))
	fmt.Printf("Content %s\n", <-ch)

	// --- Buffered channel from goroutines
	chNb := make(chan int, 2)
	go writeNumber(chNb)
	time.Sleep(1 * time.Second)
	for v := range chNb {
		fmt.Printf("Pouet %d\n", v)
		time.Sleep(1 * time.Second)
	}

	// --- Wait group
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines are finished executing")
}
