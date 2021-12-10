package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://www.google.com/",
	"https://golang.org/",
	"https://yahoo.com/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
	status   string
}

func asyncHttpGets(url string, ch chan *HttpResponse, wg *sync.WaitGroup) {
	client := http.Client{}
	if url == "https://www.google.com/" {
		time.Sleep(500 * time.Millisecond) //google is down
	}

	fmt.Printf("Fetching %s \n", url)
	resp, err := client.Get(url)
	ch <- &HttpResponse{url, resp, err, "fetched"}
	fmt.Println("sent to chan")
	wg.Done()
}

func main() {
	fmt.Println("start")
	ch := make(chan *HttpResponse, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go asyncHttpGets(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Im done")

}
