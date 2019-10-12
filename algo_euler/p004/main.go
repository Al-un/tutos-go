package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	nbAsTxt := strconv.Itoa(n)
	nbAsRune := []rune(nbAsTxt)

	length := len(nbAsRune)
	var first, last []rune

	if length%2 == 0 {
		first = nbAsRune[:length/2]
		last = nbAsRune[length/2:]
	} else {
		first = nbAsRune[:length/2]
		last = nbAsRune[(length/2 + 1):]
	}

	// https://kodify.net/go/strings/reverse-string/
	for i, j := 0, len(last)-1; i < j; i, j = i+1, j-1 {
		last[i], last[j] = last[j], last[i]
	}

	for index, val := range first {
		if val != last[index] {
			return false
		}
	}

	return true
}

func checkLargestPalindrome(upperBound int) (int, int, int) {
	a := 1
	b := 1
	largest := -1

	for i := 1; i <= upperBound; i++ {
		for j := 1; j <= upperBound; j++ {
			multiply := i * j
			if isPalindrome(multiply) && multiply > largest {
				largest = multiply
				a = i
				b = j
			}
		}
	}

	return largest, a, b
}

func main() {
	result, multiplerA, multiplerB := checkLargestPalindrome(1000)
	fmt.Println("Largest palindrome is ", result, " from ", multiplerA, " x ", multiplerB)
}
