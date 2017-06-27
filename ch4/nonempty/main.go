// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
)

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}
func main() {
	var s = []string{"hello", "", "world"}
	fmt.Println(s)
	fmt.Println(nonempty(s))

}
