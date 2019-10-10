package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Helloworld Go routine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("Counting %d\n", i)
	}
}

func alphabet() {
	for x := 'a'; x <= 'e'; x++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("Reading %c\n", x)
	}
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)

	go numbers()
	go alphabet()
	time.Sleep(3000 * time.Millisecond)

	fmt.Println("main terminated")
}
