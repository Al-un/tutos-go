package beginner

import (
	"fmt"
	"time"
)

func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		fmt.Println("tock")
	}
}

// RunTickers https://tutorialedge.net/golang/go-ticker-tutorial/
func RunTickers() {
	fmt.Println("Go Tickers tutorial!")

	go backgroundTask()

	fmt.Println("Main should be alive")
	select {}
}
