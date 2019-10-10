package main

import "fmt"

func main() {

	// finger := 4
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger")
	}
	// multiple case
	switch letter := "i"; letter {
	case "a", "u", "i", "e", "o":
		fmt.Println("Vowel")
	default:
		fmt.Println("Not a vowel")
	}
	// omitted expression
	x := 42
	switch {
	case x < 50:
		fmt.Println("Haha")
	default:
		fmt.Println("Hoho")
	}
	// fallthrough
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}
}

func number() int {
	return 15 * 5
}
