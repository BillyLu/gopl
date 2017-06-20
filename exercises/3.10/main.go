package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	for i, v := range s {
		r := (len(s) - i + 1) % 3
		if i > 0 && r == 1 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1112321111"))
}
