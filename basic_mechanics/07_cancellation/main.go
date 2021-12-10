package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cancellation()
}

func cancellation() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "signal"
		fmt.Println("send signal")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case signal := <-ch:
		fmt.Println("received signal", signal)
	case t := <-tc:
		fmt.Println("cancelled: timeout: ", t)
	}

	time.Sleep(time.Second)
	fmt.Println("--------------------")
}
