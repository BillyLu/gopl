// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, countsFiles)
	} else {
		for _, fileName := range(files) {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles)
			f.Close()
		}
	}
	
	for lineText, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, lineText)
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string][]string)  {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !arrayContains(countsFiles[text], text) {
			countsFiles[text] = append(countsFiles[text], name)
		}
	}
}


func arrayContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
