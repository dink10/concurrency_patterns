package main

import (
	"fmt"
	"time"
)

func main() {
	waitForFinished()
}

func waitForFinished() {
	ch := make(chan struct{})

	// Employee
	go func() {
		time.Sleep(500 * time.Millisecond)
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
		}
		close(ch)
		fmt.Println("employee: sent all signal")
	}()

	for v := range ch {
		fmt.Println("manager: received signal: ", v)
	}

	time.Sleep(time.Second)
}
