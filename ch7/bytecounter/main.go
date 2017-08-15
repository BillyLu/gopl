package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type byteCounter int

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

type wordCounter int

func (c *wordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	num := 0
	for scanner.Scan() {
		num++
	}
	*c += wordCounter(num)
	return num, nil
}

type lineCounter int

func (lc *lineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	num := 0
	for scanner.Scan() {
		num++
	}
	*lc += lineCounter(num)
	return num, nil
}



func main() {
	var c byteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	var wordc wordCounter
	fmt.Fprintf(&wordc, "hello, %s", name)
	fmt.Println(wordc)

	var linec lineCounter
	fmt.Fprintf(&linec, "hello, 1 \n 2 \n 3")
	fmt.Println(linec)
	fmt.Fprintf(&linec, "hello, 1 \n 2 \n 3")
	fmt.Println(linec)
}
