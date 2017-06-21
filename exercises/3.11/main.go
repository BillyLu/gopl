package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	var current, floatPoint int

	// put sign to buf and move current forward
	if strings.HasPrefix(s, "-") {
		current = 1
		buf.WriteByte('-')
	}

	// find float point index
	floatPoint = strings.Index(s, ".")
	if floatPoint == -1 {
		floatPoint = len(s)
	}

	// integer part
	for ; current < floatPoint; current++ {
		buf.WriteByte(s[current])

		step := floatPoint - current
		if step != 1 && step%3 == 1 {
			buf.WriteByte(',')
		}
	}

	// fractional part
	if floatPoint != len(s) {
		buf.WriteByte('.')
		buf.WriteString(s[current+1:])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("123"))         // 123
	fmt.Println(comma("1234"))        // 1,234
	fmt.Println(comma("12345"))       // 12,234
	fmt.Println(comma("123456"))      // 123,456
	fmt.Println(comma("1234567"))     // 1,234,567
	fmt.Println(comma("1234567.01"))  // 1,234,567.01
	fmt.Println(comma("-1234567.01")) // -1,234,567.01
}
