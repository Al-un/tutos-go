package main

import "fmt"

func main() {
	var bill = calcBill(5, 3)
	fmt.Println("bill:", bill)
	area, perim := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perim)
	area2, perim2 := rectPropsNamed(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area2, perim2)
	a, _ := rectProps(10.8, 5.6)
	fmt.Println("Area:", a)

}

func calcBill(price, qty int) int {
	return price * qty
}

func rectProps(length, width float64) (float64, float64) {
	return length * width, 2 * (length + width)
}

func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}
