package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				runtime.Gosched() // !!!!
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Nanosecond)
			}
		}()
	}

	time.Sleep(time.Millisecond)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
