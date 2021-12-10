package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: send signal")
	}()

	result := <-ch
	fmt.Println("manager: received result: ", result)

	time.Sleep(time.Second)
}
