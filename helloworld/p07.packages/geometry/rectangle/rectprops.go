package rectangle

import "math"
import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("Rectangle package initialised")
}

// Area blabla
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal Fetch the diagonal of a rectangle
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
