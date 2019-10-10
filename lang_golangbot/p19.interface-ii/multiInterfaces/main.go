package main

import "fmt"

// SalaryCalculator displays everyone salary
type SalaryCalculator interface {
	DisplaySalary()
}

// LeaveCalculator checks how much holidays is left
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

// Employee is just some dude
type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("Salary of %s %s is %d\n", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	// var s SalaryCalculator = e
	// s.DisplaySalary()
	// var l LeaveCalculator = e
	// fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	var s EmployeeOperations = e
	s.DisplaySalary()
	fmt.Println("\nLeaves left =", s.CalculateLeavesLeft())
}
