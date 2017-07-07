package main

import (
	"fmt"
)

func sum(vals ...int) (total int) {
	for _, val := range vals {
		total += val
	}
	return
}

func main() {
	fmt.Println(sum())
}
