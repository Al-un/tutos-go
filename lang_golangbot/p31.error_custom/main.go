package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

type rectAreaError struct {
	err    string
	length float64 // faulty length
	width  float64 // faulty width
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func (e *rectAreaError) Error() string {
	return e.err
}

func (e *rectAreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectAreaError) widthNegative() bool {
	return e.width < 0
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New("Area calculation failed, radius less than 0")
		// return 0, fmt.Errorf("Area calculation failed, radius %0.2f less than 0", radius)
		return 0, &areaError{"radius is negative", radius}
	}

	return math.Pi * radius * radius, nil
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err += "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}

	if err != "" {
		return 0, &rectAreaError{err, length, width}
	}

	return length * width, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		// return
	} else {
		fmt.Printf("Area of circle is %0.2f\n", area)
	}

	length, width := -5.0, -9.0
	rectArea, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}

			return
		}
		fmt.Println(err)
		return
	}

	fmt.Println("area of rect", rectArea)
}
