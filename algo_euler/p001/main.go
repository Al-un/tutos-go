package main

import "fmt"

func main() {
	sum := 0
	limit := 1000

	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			// fmt.Printf("Adding %d make it %d\n", i, sum)
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}
