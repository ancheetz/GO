package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	//infinite loop
	//annonymous function
	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)
	}

}

//this checkLink expects arguments of the link, and the channel to communicate the string through
func checkLink(link string, channel chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}

	fmt.Println(link, "is up!")
	channel <- link
}
