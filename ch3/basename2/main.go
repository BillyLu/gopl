package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slashIndex := strings.LastIndex(s, "/")
	s = s[slashIndex+1:]

	if dotIndex := strings.LastIndex(s, "."); dotIndex >= 0 {
		s = s[:dotIndex]
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}
