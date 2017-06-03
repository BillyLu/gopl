package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(index, " : ", arg)
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))
}
