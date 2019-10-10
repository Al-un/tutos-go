package main

import (
	"fmt"
)

type describer interface {
	describe()
}

type person struct {
	name string
	age  int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%s is a string\n", v)
	case int:
		fmt.Printf("%d is an int\n", v)
	case describer:
		v.describe()
	default:
		fmt.Printf("Unknown type %s: %T\n", i, i)
	}
}

func main() {
	findType("Plop")
	findType(39)
	findType(1.2)

	p := person{name: "bob", age: 43}
	findType(p)
}
