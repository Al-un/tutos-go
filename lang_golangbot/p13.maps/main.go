package main

import "fmt"

func main() {

	// personSalary := make(map[string]int)
	personSalary := map[string]int{"pouet": 123}
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	personSalary["joe"] = 0
	fmt.Println("personSalary map contents:", personSalary)
	personSalary["jamie"] = 14000
	fmt.Println("personSalary map contents:", personSalary)
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of", "bob", "is", personSalary["bob"])
	val, ok := personSalary["bob"]
	fmt.Println("Bob has a salary (", ok, "):", val)
	val, ok = personSalary["joe"]
	fmt.Println("Joe has a salary (", ok, "):", val)
	for key, val := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, val)
	}
	delete(personSalary, "mike")
	for key, val := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, val)
	}
}
