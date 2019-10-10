package main

import (
	"fmt"
)

func check(x interface{}) {
	v, ok := x.(int) // check if x is an int
	fmt.Println(v, ok)
}

func main() {
	var i interface{} = 16
	var a interface{} = "some text"
	check(i)
	check(a)
}
