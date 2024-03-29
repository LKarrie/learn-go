package concurrency

import (
	"fmt"
	"time"
)

func GoTicker() {
	// Timer 只执行一次 Ticker可以周期的执行
	ticker := time.NewTicker(time.Second)
	counter := 1
	for range ticker.C {
		fmt.Println("ticker 1")
		counter++
		if counter >= 5 {
			break
		}
	}
	ticker.Stop()

	chanInt := make(chan int)

	ticker2 := time.NewTicker(time.Second)

	go func() {
		for range ticker2.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("接收: %v\n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("sum: %v\n", sum)
			break
		}
	}

}
