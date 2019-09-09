package main

import (
	"fmt"
	"math"
)

type employee struct {
	name     string
	salary   int
	currency string
}

type rectangle struct {
	width  int
	height int
}

type circle struct {
	radius float64
}

type identity struct {
	id int
	employee
}

func (e employee) displaySalary() {
	fmt.Printf("Salary of %s is %d %s\n", e.name, e.salary, e.currency)
}

func (e employee) changeName(newName string) {
	e.name = newName
}

func (e *employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() int {
	return r.height * r.width
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	// Basic method
	emp1 := employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "US$",
	}

	emp1.displaySalary()

	// Same method name for two structures
	c := circle{
		radius: 2,
	}
	r := rectangle{
		width:  4,
		height: 2,
	}

	fmt.Printf("Circle area is %f\n", c.area())
	fmt.Printf("Rectangle area is %d\n", r.area())

	// pointer vs value
	emp1.changeName("Bob")
	emp1.changeSalary(2500)
	emp1.displaySalary()
	(&emp1).changeSalary(3500)
	emp1.displaySalary()

	// anonymous field
	i := identity{
		id:       1,
		employee: emp1,
	}
	i.displaySalary()

	// method on primitive type
	a := myInt(5)
	b := myInt(10)
	fmt.Printf("Sum is %d\n", a.add(b))
}
