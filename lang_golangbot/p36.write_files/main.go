package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

// test.txt file location is relative to the execution location
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// String
	// l, err := f.WriteString("Hello world")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// fmt.Println(l, " bytes written")

	// Bytes
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 10}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, " bytes written")

	// Strings (multiple)
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Appending
	f2, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "This is a new line and stuff"
	_, err = fmt.Fprintln(f2, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// concurrent access
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	dStatus := <-done
	if dStatus == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
