package htmlLinks

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

// Link represent a link in an HTML document.
type Link struct {
	Text string
	Href string
}

// Parse wil take in an HTML documment and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("html.Parse err in Parse fn")
		log.Fatal(err)
	}
	Dfs(doc, "")
	return nil, nil
}

func Dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Dfs(c, padding+"  ")
	}
}
