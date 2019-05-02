package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	cnt, err := reader.ReadString('\n')
	// cnt = strings.TrimSuffix(cnt, "\n")
	cnt = strings.TrimRight(string(cnt), "\r\n")

	if err != nil {
		fmt.Println("Exiting from reader...", err)
		os.Exit(1)
	}

	// count, err := strconv.ParseInt(cnt, 10, 32)
	count, err := strconv.Atoi(cnt)
	if err != nil {
		fmt.Println("Exiting from strconv...", err)
		os.Exit(1)
	}

	var words []string

	for i := int(0); i < count; i++ {
		// text, err := reader.ReadString('\n')
		text, _, err := reader.ReadLine()
		// text = strings.TrimSuffix(text, "\n")
		// text = strings.TrimRight(string(text), "\r\n")
		if err != nil {
			fmt.Println("Exiting from for...", err)
			os.Exit(1)
		}
		words = append(words, string(text))
	}

	for _, word := range words {
		// fmt.Println("=> ", word)
		var even []byte
		var odd []byte

		for i, w := range []byte(word) {
			// fmt.Println("=> ", string(w))
			if i%2 == 0 {
				even = append(even, w)

			} else {
				odd = append(odd, w)
			}
		}
		fmt.Println(string(even), string(odd))
	}
}
