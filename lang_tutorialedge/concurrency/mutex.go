package concurrency

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

func deposit(amount int, done chan bool) {
	// mutex.Lock()

	fmt.Printf("Deposing %d to balance %d\n", amount, balance)
	balance += amount
	// mutex.Unlock()

	done <- true
}

func withdraw(amount int, done chan bool) {
	// mutex.Lock()

	fmt.Printf("Withdrawing %d from balance %d\n", amount, balance)
	balance -= amount
	// mutex.Unlock()

	done <- true
}

// RunMutex https://tutorialedge.net/golang/go-mutex-tutorial/
func RunMutex() {
	fmt.Printf("Start balance before all operations: %d\n", balance)

	done := make(chan bool)
	go withdraw(700, done)
	go deposit(500, done)
	<-done
	<-done

	fmt.Printf("Total balance after all operations: %d\n", balance)
}
