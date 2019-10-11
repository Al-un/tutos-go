package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "From server 1"
}

func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "From server 2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	// Learn about select
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	// test random selection
	time.Sleep(3 * time.Second)

	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}

	// Use-case
	// chProcess := make(chan string)
	// go process(chProcess)
	// for {
	// 	time.Sleep(1000 * time.Millisecond)
	// 	select {
	// 	case v := <-chProcess:
	// 		fmt.Println("Received ", v)
	// 	default:
	// 		fmt.Println("Nothing received")
	// 	}
	// }

	// Deadlock
	// nil channel
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("Should not received ", v)
	default:
		fmt.Println("Should only select default")
	}
}
