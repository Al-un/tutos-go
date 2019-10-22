package beginner

import (
	"fmt"
	"sort"
)

type programmer struct {
	Age int
}

type byAge []programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// RunSort : https://tutorialedge.net/golang/go-sorting-with-sort-tutorial/
func RunSort() {
	arr := []int{4, 2, 6, 3, 7}
	fmt.Println("Unsorted array: ", arr)

	sort.Ints(arr)
	fmt.Println("Sorted array: ", arr)

	programmers := []programmer{
		programmer{Age: 30},
		programmer{Age: 20},
		programmer{Age: 50},
		programmer{Age: 1000},
	}
	fmt.Println("Unsorted programmers: ", programmers)

	sort.Sort(byAge(programmers))

	fmt.Println("Sorted programmers: ", programmers)
}
