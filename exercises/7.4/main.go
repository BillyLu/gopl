package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return 0, err
}

func newReader(s string) io.Reader {
	return &stringReader{s}
}

func main() {
	r := newReader("<html></html>")
	doc, err := html.Parse(r)
	fmt.Printf("%#v \n%#v\n", doc, err)
}
