package main

import "fmt"

type vowelFinder interface {
	findVowels() []rune
}

type myString string

func (ms myString) findVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func main() {
	name := myString("Sam Anderson")
	var v vowelFinder
	v = name
	fmt.Printf("Vowels of %s are %c\n", name, v.findVowels())
}
