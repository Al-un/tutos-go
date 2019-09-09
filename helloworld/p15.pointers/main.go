package main

import "fmt"

func main() {

	b := 255
	var a = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("Address of b is", a)
	fmt.Println("Value of b is", *a)
	*a++
	fmt.Println("Value of b is", b)

	// nil pointer
	c := "pouet"
	var d *string
	fmt.Println("d is", d)
	d = &c
	fmt.Println("d is", d)
	fmt.Println("c via *d is", *d)

	// pointer & functions
	changePointer(a)
	fmt.Println("Value of b after changePointer is", b)

}

func changePointer(val *int) {
	*val = 55
}
