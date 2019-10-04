package main

import "fmt"

type describer interface {
	Describe()
}

func main() {
	var d1 describer
	// d1.Describe()
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T and value %v\n", d1, d1)
	}
}
