package main

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}
	wg.Wait()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithChannel(&wg, ch)
	}
	wg.Wait()

	fmt.Println("Final value of x: ", x)
	fmt.Println("Final value of y: ", y)
}
