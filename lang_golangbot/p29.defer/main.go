package main

import (
	"fmt"
	"sync"
)

type person struct {
	first string
	last  string
}

type rect struct {
	width  int
	height int
}

type rectArea struct {
	area   int
	source rect
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.first, p.last)
}

func finished(pouet string) {
	fmt.Printf("Finished! %s\n", pouet)
}

func largest(nums []int) {

	fmt.Println("Starting")
	max := nums[0]
	defer finished(fmt.Sprintf("Poueeeeeeeeet: %d", max))

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Printf("Largest number of %v is %d\n", nums, max)
}

func printA(prefix string, a int) {
	fmt.Println(prefix, " a is ", a)
}

func reveseWithDefer(text string) {
	fmt.Println("Origin input is ", text)
	fmt.Println("Reversing: ")

	defer fmt.Printf("\n")
	for _, v := range []rune(text) {
		defer fmt.Printf("%c", v)
	}
}

func (r rect) area(wg *sync.WaitGroup, ch chan rectArea) {
	defer wg.Done()

	if r.width < 0 {
		fmt.Printf("Invalid width for rect %v\n", r)
		ch <- rectArea{area: -1, source: r}
		return
	}

	if r.height < 0 {
		fmt.Printf("Invalid height for rect %v\n", r)
		ch <- rectArea{area: -1, source: r}
		return
	}

	area := r.width * r.height
	ch <- rectArea{area: area, source: r}
	fmt.Printf("React %v has area %d\n", r, area)
}

func main() {
	// simple defer
	p := person{"aa", "bb"}
	defer p.fullName()

	nums := []int{43, 235, 32432, 23, 32, 3542}
	largest(nums)

	// defer in struct
	a := 5
	printA("stardard", a)
	defer printA("deferred", a)
	a = 10
	printA("end", a)

	// defer in methods
	reveseWithDefer("HelloWorld感じ")

	// defer use case with WaitGroup
	r1 := rect{width: -2, height: 2}
	r2 := rect{width: 2, height: -2}
	r3 := rect{width: 2, height: 2}
	r4 := rect{width: 2, height: 4}
	r5 := rect{width: 2, height: 3}
	r6 := rect{width: 3, height: 3}
	reactCh := make(chan rectArea, 12)

	var wg sync.WaitGroup
	rects := []rect{r1, r2, r3, r4, r5, r6}
	for _, r := range rects {
		wg.Add(1)
		go r.area(&wg, reactCh)
	}
	// -- trying to combine with channels
	// for v := range reactCh {
	// 	fmt.Printf("Getter area %d for rect %v from channel\n", v.area, v.source)
	// }
	wg.Wait()
}
