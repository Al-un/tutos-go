package internalRepresentation

import (
	"fmt"
)

type Tester interface {
	Test()
}

type MyFloat float64

func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func toRenameMain() {
	var t Tester
	f := MyFloat(89.7)
	t = f

	describe(t)
	t.Test()
}
