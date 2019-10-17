package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	defer fmt.Println("Deferred call normally from fullName")

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("Return normally from fullName")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
		debug.PrintStack()
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("Oh no! B panicked")
}

func c() {
	defer recovery()
	c := []int{5, 7, 3}
	fmt.Println(c[3])
	fmt.Println("Returned normally from c")
}

func main() {
	defer fmt.Println("Deferre from main")
	first := "Pouet"
	fullName(&first, nil)
	fmt.Println("[1] Returned from main")

	a()
	fmt.Println("[2] Returned from main")

	c()
	fmt.Println("[3] Returned from main")
}
