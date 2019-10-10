package main

import "fmt"

func main() {

	find(89, 90, 89, 95)
	find(45, 35, 24, 65489516, 45, 626626)
	find(456, 123, 987, 654, 321, 654)
	nums := []int{89, 80, 95}
	find(89, nums...)

	// modify a slice
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)

}

func find(num int, nums ...int) {
	fmt.Printf("Type if nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "fond at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
}
