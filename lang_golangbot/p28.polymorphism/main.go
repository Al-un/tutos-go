package main

import "fmt"

// Income for money!!!
type Income interface {
	calculate() int
	source() string
}

type fixedBilling struct {
	projectName  string
	biddedAmount int
}

type timeAndMaterial struct {
	projectName string
	nbOfHours   int
	hourlyRate  int
}

type ads struct {
	adName     string
	cpc        int // cost per click
	nbOfClicks int
}

func (fb fixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb fixedBilling) source() string {
	return fb.projectName
}

func (tm timeAndMaterial) calculate() int {
	return tm.nbOfHours * tm.hourlyRate
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}

func (a ads) calculate() int {
	return a.cpc * a.nbOfClicks
}

func (a ads) source() string {
	return a.adName
}

func calculateNetIncome(ic []Income) {
	var netIncome = 0

	for _, income := range ic {
		fmt.Printf("%s has income %d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}

	fmt.Printf("Total income is %d\n", netIncome)
}

func main() {
	p1 := fixedBilling{projectName: "prout", biddedAmount: 1000}
	p2 := timeAndMaterial{projectName: "Pouet", nbOfHours: 10, hourlyRate: 8}
	p3 := timeAndMaterial{projectName: "Pouet", nbOfHours: 10, hourlyRate: 12}
	p4 := ads{adName: "plop", cpc: 10, nbOfClicks: 20}
	incomes := []Income{p1, p2, p3, p4}

	calculateNetIncome(incomes)
}
