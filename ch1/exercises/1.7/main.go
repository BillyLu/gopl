// 1.7 - 1.9
package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"strings"
)


func fillURL(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url 
}

func main()  {
	for _, url := range(os.Args[1:]) {
		url = fillURL(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("http status %d\n", resp.StatusCode)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}