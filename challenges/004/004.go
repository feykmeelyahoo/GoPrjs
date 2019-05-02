package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	// var arr []int32
	m := make(map[string]string)

	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(readLine(reader), " ")
		m[arrTemp[0]] = arrTemp[1]
	}

	var sl []string

	for i := 0; i < int(n); i++ {
		line := readLine(reader)
		sl = append(sl, line)

	}

	for _, k := range sl {
		_, found := m[k]

		if found {
			fmt.Printf("%s=%s\n", k, m[k])
		} else {
			fmt.Printf("Not found\n")
		}

		// fmt.Println(k)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
