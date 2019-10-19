package intermediate

import "fmt"

func subVariadic(args ...interface{}) {
	fmt.Println("Subvariadic args are: ", args)
}

func myVariadicFunction(args ...string) {
	fmt.Println("Args are ", args)
	subVariadic(args)
}

// RunVariadicFunction https://tutorialedge.net/golang/go-variadic-function-tutorial/
func RunVariadicFunction() {
	myVariadicFunction("hello", "world", "!")
}
