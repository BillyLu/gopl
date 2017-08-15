package main

import "io"
import "os"
import "fmt"

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}

// CountingWriter returns a new Writer that wraps the original,
//and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.written
}

func main() {
	cw, written := CountingWriter(os.Stdout)
	fmt.Fprint(cw, "hello\n")
	fmt.Println(*written)
	fmt.Fprint(cw, "world\n")
	fmt.Println(*written)
}
