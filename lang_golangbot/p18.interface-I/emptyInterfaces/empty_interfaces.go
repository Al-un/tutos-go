package emptyInterface

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func toMakeMain() {
	s := "Helloword"
	describe(s)
	i := 55
	describe(i)

	strt := struct {
		name string
	}{
		name: "Bob",
	}
	describe(strt)
}
