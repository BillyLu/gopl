// reverse reverses a slice of ints in place.
package main

import (
	"fmt"
)

func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{1, 2, 3: 4, 4, 5}
	fmt.Println("before ", a)
	reverse(&a)
	fmt.Println("after ", a)

	var runes []rune
	fmt.Println(runes == nil, len(runes) == 0)
}
