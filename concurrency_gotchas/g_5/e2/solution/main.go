package main

func main() {
	var a []int = nil
	c := make(chan struct{})

	go func() {
		a = make([]int, 3)
		c <- struct{}{}
	}()

	<-c
	// The next line will not panic for sure.
	a[0], a[1], a[2] = 0, 1, 2
}
