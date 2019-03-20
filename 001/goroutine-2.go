package main

import (
	"fmt"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.stackoverflow.com",
		"http://www.facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checklink2(link, c)
		fmt.Println("in for loop !!!")
	}

	for {
		checklink2(<-c, c)
	}

	fmt.Println("Exit!!!")
}

func checklink2(url string, c chan string, tSleep time.Duration=0) {
	fmt.Println("checklink2 START", url)
	time.Sleep(tSleep)
	fmt.Println("checklink2 END ", url)
	c <- url
}
