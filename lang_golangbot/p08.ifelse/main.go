package main

import "fmt"

func main() {
	// num := 10
	if num := 10; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
