package main

import (
	"fmt"
)

func runPipe(id string, i chan int, o chan int) {
	defer close(o)

	for x := range i {
		fmt.Printf("goroutine '%s': %d\n", id, x)
		o <- x + 1
	}

	fmt.Println("done")
}

func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go runPipe("one", c0, c1)
	go runPipe("two", c1, c2)
	go runPipe("three", c2, c3)

	go func() {
		c0 <- 1
		c0 <- 10
		c0 <- 100
		c0 <- 1000
		c0 <- 10000
		c0 <- 100000
		fmt.Println("Sent all numbers to c0")
		close(c0)
	}()

	fmt.Println("Sent all numbers to c0")

	for x := range c3 {
		fmt.Printf("out: %d\n", x)
	}
}
