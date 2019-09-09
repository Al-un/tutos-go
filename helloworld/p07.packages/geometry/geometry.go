package main

import (
	"fmt"
	"log"

	"github.com/al-un/tutos/helloworld/p07.packages/geometry/rectangle"
)

// 1. Variables
var _ = rectangle.Area // error silencer
var rectLen, rectwidth float64 = -6, 7

// 2. init
func init() {
	println("main package initialisation")
	if rectLen < 0 {
		log.Fatal("length < 0")
	}
	if rectwidth < 0 {
		log.Fatal("width < 0")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")

	fmt.Printf("Area is %.2f\n", rectangle.Area(rectLen, rectwidth))
	fmt.Printf("Diagonal is %.2f\n", rectangle.Diagonal(rectLen, rectwidth))
}
