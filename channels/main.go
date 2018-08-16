package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.in",
		"http://google.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	go makeRequest(<-c)
	// }

	// for {
	// 	go makeRequest(<-c, c)
	// }

	// for l := range c {
	//
	// 	go makeRequest(l, c) //Same as above for infinite
	// }

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			makeRequest(l, c)
		}(l)
	}

	fmt.Println("Done")
}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Link might be down ", link)
		c <- link
		return
	}

	fmt.Println("Link is working ", link)
	c <- link

}
