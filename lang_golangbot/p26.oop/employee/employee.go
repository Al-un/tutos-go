package employee

import "fmt"

// Employee is pouet
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// New generates an `Employee`
func New(first string, last string, leaves int, takenLeaves int) employee {
	e := employee{firstName: first, lastName: last, totalLeaves: leaves, leavesTaken: takenLeaves}
	return e
}

// LeavesRemaining displays how many holidays left
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, e.totalLeaves-e.leavesTaken)
}
