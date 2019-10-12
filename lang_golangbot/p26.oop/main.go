package main

import (
	"github.com/al-un/tutos-go/lang_golangbot/p26.oop/employee"
)

func main() {
	// e := employee.Employee{
	// 	FirstName:   "Bob",
	// 	LastName:    "Jack",
	// 	LeavesTaken: 3,
	// 	TotalLeaves: 17,
	// }
	// e.LeavesRemaining()

	e2 := employee.New("Pouet", "Plop", 12, 2)
	e2.LeavesRemaining()
}
