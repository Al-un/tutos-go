package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	age = 42
	fmt.Println("My age is", age)
	age = 29
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My age is", age)

	// var width, height int = 100, 50
	var width, height = 100, 50
	fmt.Println("width is", width, "height is", height)

	// Global declaration
	var (
		name    = "Pouet"
		address = "trou du cul du monde"
		partner string
	)
	fmt.Println("Name", name, "is living", address, "with", partner)
	partner = "Plop"
	fmt.Println("Name", name, "is living", address, "with", partner)

	// Shorthand
	firstName, lastName := "toto", "Pouet"
	fmt.Println("Fullname is", firstName, lastName)
	firstName, middleName := "Toto", "De"
	fmt.Println("Fullname is", firstName, middleName, lastName)

	// Math
	a, b := 7.2, 5.0
	c := math.Min(a, b)
	fmt.Println("Min is", c)

}
