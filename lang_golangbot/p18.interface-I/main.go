package main

import "fmt"

type vowelFinder interface {
	findVowels() []rune
}

type myString string

func (ms myString) findVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

// Second part of the tuto

// SalaryCalculator Must calculate salary
type SalaryCalculator interface {
	CalculateSalary() int
}

// Permanent is permanent employee
type Permanent struct {
	empID    int
	basicPay int
	pf       int
}

// Contract can be anything?

type Contract struct {
	empID    int
	basicPay int
}

// CalculateSalary Pouet pouet
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// CalculateSalary Plop plop
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) int {
	expense := 0

	for _, v := range s {
		expense += v.CalculateSalary()
	}

	return expense
}

func main() {
	name := myString("Sam Anderson")
	var v vowelFinder
	v = name
	fmt.Printf("Vowels of %s are %c\n", name, v.findVowels())

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExp := totalExpense(employees)

	fmt.Printf("Total salary is %d\n", totalExp)
}
