package main

import "fmt"

type describer interface {
	Describe()
}

type person struct {
	name string
	age  int
}

func (p person) Describe() {
	fmt.Printf("%s has %d years old\n", p.name, p.age)
}

type address struct {
	state   string
	country string
}

func (a *address) Describe() {
	fmt.Printf("State %s is in country %s\n", a.state, a.country)
}

func main() {
	var d1 describer
	p1 := person{"Bob", 42}
	d1 = p1
	d1.Describe()
	p2 := person{"Jack", 39}
	d1 = &p2
	d1.Describe()

	var d2 describer
	a1 := address{"Washington", "USA"}
	d2 = &a1

	d2.Describe()
}
