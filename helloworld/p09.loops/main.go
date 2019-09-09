package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\n")

	// Shorthand
	j := 0
	for j <= 10 {
		fmt.Printf("%d ", j)
		j += 2
	}
	fmt.Printf("\n")

	// Double loop?
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
