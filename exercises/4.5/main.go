//Exercis e 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice
package main

import (
	"fmt"
)

func removeDup(strings []string) []string {
	var w int
	for _, s := range strings {
		if s == strings[w] {
			continue
		}
		w++
		strings[w] = s
	}

	return strings[:w+1]
}

func main() {
	var strings = []string{"a", "b", "c", "c", "c", "d"}
	strings = removeDup(strings)
	fmt.Println(strings)
}
