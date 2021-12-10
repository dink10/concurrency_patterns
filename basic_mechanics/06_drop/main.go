package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

func drop() {
	const capacity = 5
	ch := make(chan string, capacity)

	go func() {
		for signal := range ch {
			fmt.Println("received signal: ", signal)
		}
	}()

	const work = 20
	for i := 0; i < work; i++ {
		select {
		case ch <- fmt.Sprintf("signal #%d", i):
			fmt.Println("send signal #", i)
		default:
			fmt.Println("skipped signal")
		}
	}

	close(ch)

	time.Sleep(5 * time.Second)
}
