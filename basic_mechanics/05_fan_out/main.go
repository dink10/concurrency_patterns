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

	for e := 1; e <= emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "finished task by employee: " + fmt.Sprint(emp)
		}(e)
	}

	for emps > 0 {
		res := <-ch
		fmt.Println("manager received result: ", res)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("------------------")
}
