package main

import "fmt"

func main() {

	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	// auto length
	a2 := [...]int{12, 78}
	// a2 := [3]int{12}
	// a2 := [3]int{12, 78, 50}
	fmt.Println(a2)

	// value type
	b := [...]string{"USA", "France", "UK"}
	c := b
	b[0] = "Singapore"
	fmt.Println(b)
	fmt.Println(c)

	// local vs in function
	num := [...]int{1, 2, 3}
	fmt.Println("Before calling changeLocal", num)
	changeLocal(num)
	fmt.Println("After calling changeLocal", num)

	// length
	x := [...]int{1, 2, 3, 4}
	fmt.Println("Length of x:", len(x))
	for i := 0; i < len(x); i++ {
		fmt.Printf("%d index of x is %d\n", i, x[i])
	}

	// range
	for i, v := range a {
		fmt.Printf("%d index of a is %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf(" %d\n", v)
	}

	// multi dimension array
	animal := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	var brands [3][2]string
	brands[0][0] = "Apple"
	brands[0][1] = "Samsung"
	brands[1][0] = "Microsoft"
	brands[1][1] = "Google"
	brands[2][0] = "AT&T"
	brands[2][1] = "T-Mobile"
	printArray(animal)
	fmt.Printf("\n")
	printArray(brands)

	// Slices
	d := [5]int{76, 77, 78, 79, 80}
	var e = d[1:4]
	// var e []int = d[1:4]
	fmt.Println(e)
	f := []int{6, 7, 8}
	fmt.Println(f)

	// Operate
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for dsliceI := range dslice {
		dslice[dsliceI]++
	}
	fmt.Println("array after", darr)

	// slices impact on array
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	// len vs cap
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	// make
	pouet := make([]int, 5, 5)
	fmt.Println(pouet)

	// under the hood
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "Renault", "Peugeot", "Citroen", "BMW")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	// nil?
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// concat
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// pass to functions
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	// multi dimensional slices
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	// memory management
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func changeLocal(num [3]int) {
	num[0] = 42
	fmt.Println("Inside changeLocal func", num)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
