package main

import "fmt"

func main() {
	const hello = "Hello Go!"
	var name = "Sam"
	fmt.Printf("type %T value %v\n", name, name)
	fmt.Printf("type %T value %v", hello, hello)

	// --- type mixing
	// var defaultName = "Sam" //allowed
	// type myString string
	// var customName myString = "Sam" //allowed
	// customName = defaultName        //not allowed

	const a = 5
	var intVar = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
