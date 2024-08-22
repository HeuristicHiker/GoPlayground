package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	preUrl := "http://"
	links := []string{
		preUrl + "google.com",
		preUrl + "facebook.com",
		preUrl + "stackoverflow.com",
		preUrl + "golang.org",
		preUrl + "amazon.com",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// using range of channel
	// runs every time channel produces return
	for l := range c {
		// blocking each time
		// sleep here just sleeps after completion
		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		// passed l to make sure links are copied and updated
	}
}

func checkLink(link string, c chan string) {
	// sleep here just means it waits to start the fetch
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
