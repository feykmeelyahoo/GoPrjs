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
	var count time.Duration = 4

	for _, link := range links {
		count--
		sTime := count
		go checklink(link, c, count*time.Second)
		fmt.Println("in for loop ", link, sTime, " !!!")
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Exit!!!")
}

func checklink(url string, c chan string, tSleep time.Duration) {
	fmt.Println("checklink START", url)
	str := fmt.Sprintf("%s is ok", url)

	time.Sleep(tSleep)
	fmt.Println("checklink END ", url)

	c <- str
}
