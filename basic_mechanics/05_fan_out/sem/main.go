package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanOut()
}

func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	const capacity = 5
	sem := make(chan struct{}, capacity)

	for e := 1; e <= emps; e++ {
		go func(emp int) {
			sem <- struct{}{}
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "finished task by employee: " + fmt.Sprint(emp)
			}
			<-sem
		}(e)
		time.Sleep(500 * time.Millisecond)
	}

	for emps > 0 {
		res := <-ch
		fmt.Println("manager received result: ", res)
		emps--
	}

	time.Sleep(20 * time.Second)
}
