package main

import (
	"fmt"
)

var pc [8]byte

func init() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
}
func main() {
	fmt.Println(pc)
}
