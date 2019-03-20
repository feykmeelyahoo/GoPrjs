package main

import (
	"fmt"
	"net/http"
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
		fmt.Println("in for loop !!!")
		go checklink2(link, c)
	}

	for {
		str := <-c
		fmt.Println("Second for loop", str)

		go checklink2(str, c)
		// fmt.Println(<-c)

	}

	fmt.Println("Exit!!!")
}

func checklink2(url string, c chan string) {
	_, err := http.Get(url)
	fmt.Println("Try to get", url)

	if err != nil {
		fmt.Println(url, "might be down !!!")
		c <- url
		return
	}
	fmt.Println(url, " is up ...")
	c <- url
	time.Sleep(2 * time.Second)
}
