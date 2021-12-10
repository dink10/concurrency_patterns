package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var numbers []int // nil

	// start a goroutine to initialise array
	go func() {
		numbers = make([]int, 2)
	}()

	// do something synchronous
	if numbers == nil {
		time.Sleep(time.Second)
	}
	numbers[0] = 1 // will sometimes panic here
	fmt.Println(numbers[0])
}
