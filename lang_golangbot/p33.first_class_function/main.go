package main

import "fmt"

type mathOperation func(a int, b int) int

type student struct {
	first   string
	last    string
	grade   string
	country string
}

func simpleHigherOrderFunc(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simpleFunc() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}

	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	// func as value
	a := func() {
		fmt.Println("Hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	// inline func
	func(n string) {
		fmt.Println("Another anonymous function with an argument: ", n)
	}("Gogogo!")

	// Typing
	var sum mathOperation = func(a int, b int) int {
		return a + b
	}
	var mult mathOperation = func(a int, b int) int {
		return a * b
	}
	s := sum(5, 6)
	m := mult(5, 6)
	fmt.Printf("Sum is %d and Mult it %d\n", s, m)

	// Higher order function
	f := func(a, b int) int {
		return a + b
	}
	simpleHigherOrderFunc(f)

	// Return a func from a func
	simpleF := simpleFunc()
	fmt.Println(simpleF(60, 7))

	// Closure
	aa := 5
	func() {
		fmt.Println("aa = ", aa)
	}()

	// Closure 2
	aaa := appendStr()
	bbb := appendStr()
	fmt.Println(aaa("World"))
	fmt.Println(bbb("Everyone"))
	fmt.Println(aaa("Gopher"))
	fmt.Println(bbb("!"))

	// Closure usage
	s1 := student{
		first:   "bob",
		last:    "joe",
		grade:   "A",
		country: "Plop",
	}
	s2 := student{
		first:   "bobby",
		last:    "joey",
		grade:   "B",
		country: "Ploppy",
	}
	students := []student{s1, s2}
	filteredStudents := filter(students, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(filteredStudents)

	filterByCountryStudents := filter(students, func(s student) bool {
		if s.country == "Plop" {
			return true
		}
		return false
	})
	fmt.Println(filterByCountryStudents)

	// map implementation
	arr := []int{5, 22, 8, 2}
	r := iMap(arr, func(n int) int { return n * 5 })
	fmt.Println(r)
}
