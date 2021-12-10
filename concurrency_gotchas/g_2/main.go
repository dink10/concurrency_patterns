package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://yahoo.com/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
	status   string
}

func asyncHttpGets(url string, ch chan *HttpResponse) {
	client := http.Client{}
	if url == "http://www.google.com/" {
		time.Sleep(500 * time.Millisecond) //google is down
	}

	fmt.Printf("Fetching %s \n", url)
	resp, err := client.Get(url)
	ch <- &HttpResponse{url, resp, err, "fetched"}
	fmt.Println("sent to chan")
}

func main() {
	fmt.Println("start")
	ch := make(chan *HttpResponse, len(urls))
	for _, url := range urls {
		go asyncHttpGets(url, ch)
	}

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Im done")

}
