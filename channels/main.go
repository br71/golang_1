package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// checkLink(link)
	// Go routine
	// Receiving message is blocking event
	// fmt.Println(<-c)
	//fmt.Println(<-c)

	// Receive all messages from goroutines
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	// repeating goroutine, below is short syntax (infinite loop)
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }

	// Infinite loop
	// for {
	//	go checkLink(<-c, c)
	//}

	// alternate syntax for above loop
	for l := range c {
		// With this main rutine is paused
		// Main rutine must be livee all the time
		//time.Sleep(5 * time.Second)
		// Go routine prefetch same link

		// go checkLink(l, c)

		// Function literal = anonimous finction, lamda
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// Be go rutina
// func checkLink(link string)

func checkLink(link string, c chan string) {

	// Second sleep option
	// time.Sleep(5 * time.Second)

	// Blockig function call
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		// c <- "Migt be down"
		c <- link
		return
	}

	fmt.Println(link, "is up")
	// c <- "It is up"
	c <- link
}
