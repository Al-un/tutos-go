package main

import (
	"fmt"

	"github.com/al-un/tutos-go/lang_golangbot/p16.struct/computer"
)

// Address ...
type Address struct {
	city, state string
}

// Employee ...
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Person ...
type Person struct {
	string
	int
	Address
}

type SomeOne struct {
	name    string
	address Address
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	var emp4 Employee
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	var emp7 Employee
	emp7.firstName = "Jack"
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println("Employee 3", emp3)
	fmt.Println("Employee 4", emp4)
	fmt.Println("Employee 5", emp5)
	fmt.Println("Employee 1 First:", emp1.firstName)
	fmt.Println("Employee 1 Last:", emp1.lastName)
	fmt.Println("Employee 1 Age:", emp1.age)
	fmt.Printf("Employee 1 Salary: %d $\n", emp1.salary)
	fmt.Println("Employee 7", emp7)
	fmt.Println("emp8: ", emp8)
	fmt.Println("*emp8: ", *emp8)
	fmt.Println("First Name: ", (*emp8).firstName)
	fmt.Println("First Name: ", emp8.firstName)
	fmt.Println("Age: ", (*emp8).age)

	p := Person{"Naveen", 50, Address{city: "city", state: "state"}}
	fmt.Println(p)
	p.string = "Pouet"
	p.int = 45
	p.Address.state = "STATE"
	fmt.Println(p)

	so := SomeOne{"Dude", Address{city: "Chicago", state: "Illinois"}}
	fmt.Println(so)
	so.address = Address{
		city: "New York",
	}
	fmt.Println(so)

	var spec computer.Spec
	spec.Maker = "Apple"
	spec.Price = 5000
	fmt.Println("Spec: ", spec)

	add1 := Address{"city", "state"}
	add2 := Address{"city", "state"}
	fmt.Println("Equals: ", add1 == add2)

}
