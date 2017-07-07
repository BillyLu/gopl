// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func getHtmlDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("geting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s: %s", url, resp.Status)
	}

	return doc, nil
}

func populate(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	if n.FirstChild != nil {
		populate(count, n.FirstChild)
	}
	if n.NextSibling != nil {
		populate(count, n.NextSibling)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getHtmlDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail : %v\n", err)
			os.Exit(1)
		}
		count := make(map[string]int)
		populate(count, doc)
		for ele, num := range count {
			fmt.Printf("element %s: %d \n", ele, num)
		}
	}
}
