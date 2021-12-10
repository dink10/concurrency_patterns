package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	pooling()
}

func pooling() {
	ch := make(chan string) // can be buffered like ch := make(chan string, 3)
	var wg sync.WaitGroup
	const emps = 5
	wg.Add(emps)
	for i := 1; i <= emps; i++ {
		go func(emp int) {
			defer wg.Done()
			for task := range ch {
				fmt.Printf("employee %d : received task %s\n", emp, task)
			}
			fmt.Printf("employee %d : received shutdown task\n", emp)
		}(i)
	}

	const work = 10
	for i := 1; i <= work; i++ {
		ch <- "paper:" + strconv.Itoa(i)
		fmt.Println("manager: send task")
	}
	close(ch)

	wg.Wait()
}
