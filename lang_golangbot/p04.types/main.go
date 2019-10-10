package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := true, false
	c, d := a && b, a || b
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)

	// type & size
	var e = 89
	f := 95
	fmt.Println("e:", e, "f:", f)
	fmt.Printf("e (%T) size: %d\nf (%T) size: %d", e, unsafe.Sizeof(e), f, unsafe.Sizeof(f))

	// integer / floating
	g, h := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", f, g)
	sum := g + h
	diff := g - h
	fmt.Println("sum", sum, "diff", diff)
	i, j := 56, 89
	fmt.Println("sum", i+j, "diff", i-j)

	// complex
	k, l := 6+7i, complex(6, 7)
	cadd := k + l
	fmt.Println("sum:", cadd)
	cmul := k * l
	fmt.Println("product:", cmul)

	// string
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	// type conversion
	m, n := 55, 67.8
	o := m + int(n)
	p := float64(m) + n
	fmt.Println("o:", o, "p:", p)
}
