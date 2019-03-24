package main

import (
	"fmt"
	"time"
)

func myFunc(seconds time.Duration) {
	fmt.Println("Hello World")
	time.Sleep(seconds * time.Second)
}

func coolFunc(a func(time.Duration), sec time.Duration) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a(sec)
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	coolFunc(myFunc, 2)
}
