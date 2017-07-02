//Exercis e 4.9: Wr ite a program wordfreq to rep ort the fre quency of each word in an inp ut text
//file. Cal l input.Split(bufio.ScanWords) before the firs t call to Scan to bre ak the inp ut int o
//word s instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	for w, n := range counts {
		fmt.Printf("%q\t%d\n", w, n)
	}

	if err := input.Err; err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

}
