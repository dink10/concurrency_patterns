package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	WaitForTask()
}

func WaitForTask() {
	ch := make(chan string)

	// Employee
	go func() {
		p := <-ch
		fmt.Println("employee : received task : ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"
	fmt.Println("manager : send task")

	time.Sleep(time.Second)
}
