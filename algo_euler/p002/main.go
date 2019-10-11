package main

import "fmt"

func main() {
	sum := 2
	upperBound := 4000000
	// upperBound := 10
	fiboBuffer := []int{1, 2}
	fiboValue := 0

	for fiboBuffer[0]+fiboBuffer[1] <= upperBound {
		// calculate FiboValue
		fiboValue = fiboBuffer[0] + fiboBuffer[1]

		// Update buffer
		fiboBuffer[0] = fiboBuffer[1]
		fiboBuffer[1] = fiboValue

		// check for sum
		if fiboValue%2 == 0 {
			sum += fiboValue
			fmt.Printf("Adding %d making it %d\n", fiboValue, sum)
		}
	}

	fmt.Printf("Sum is %d and last value is %d\n", sum, fiboValue)
}
