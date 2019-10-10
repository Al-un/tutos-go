package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Pouet")
	time.Sleep(1 * time.Second)
	fmt.Println("Plop")
	done <- true
}

func digitProducer(number int, digitCh chan int) {
	for number != 0 {
		digit := number % 10
		number /= 10
		digitCh <- digit
	}

	close(digitCh)
}

func calcSquare(number int, squareop chan int) {
	sum := 0

	digitCh := make(chan int)
	go digitProducer(number, digitCh)
	for digit := range digitCh {
		sum += digit * digit
	}
	// for number != 0 {
	// 	digit := number % 10
	// 	sum += digit * digit
	// 	number /= 10
	// }

	squareop <- sum
}

func calcCube(number int, cubeop chan int) {
	sum := 0

	digitCh := make(chan int)
	go digitProducer(number, digitCh)
	for digit := range digitCh {
		sum += digit * digit * digit
	}

	// for number != 0 {
	// 	digit := number % 10
	// 	sum += digit * digit * digit
	// 	number /= 10
	// }

	cubeop <- sum
}

func sendData(sendCh chan<- int) {
	sendCh <- 10
}

func producer(chnl chan int) {
	for i := 0; i <= 10; i++ {
		chnl <- i
	}

	close(chnl)
}

func main() {
	// ----- Simple channel
	// var a chan int
	// if a == nil {
	// 	fmt.Println("a chan is nil. Defining it")
	// 	a = make(chan int)
	// 	fmt.Printf("a chan is type %T\n", a)
	// }

	done := make(chan bool)
	fmt.Println("Going to go-routine")
	go hello(done)
	<-done

	// ----- Interaction!
	number := 589
	sqrCh := make(chan int)
	cubeCh := make(chan int)
	go calcSquare(number, sqrCh)
	go calcCube(number, cubeCh)

	square, cube := <-sqrCh, <-cubeCh
	fmt.Printf("Square %d + Cube %d = %d\n", square, cube, square+cube)

	// // ----- Deadlock
	// lockCh := make(chan int)
	// lockCh <- 5

	// ----- Unidirectional channel
	sendCh := make(chan int)
	go sendData(sendCh)
	fmt.Println(<-sendCh)

	// ----- Channel closing
	chnl := make(chan int)
	go producer(chnl)
	for {
		v, ok := <-chnl
		if ok == false {
			fmt.Println("End detected, Finished!")
			break
		}

		fmt.Printf("Received from producer %d\n", v)
	}

	chnl2 := make(chan int)
	go producer(chnl2)
	for v := range chnl2 {
		fmt.Printf("Hey %d\n", v)
	}

	fmt.Println("main terminated")
}
