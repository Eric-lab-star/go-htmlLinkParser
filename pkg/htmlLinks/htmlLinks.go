package htmlLinks

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// Link represent a link in an HTML document.
type Link struct {
	Text string
	Href string
}

// Parse wil take in an HTML documment and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("html.Parse err in Parse fn")
		log.Fatal(err)
	}
	links := []Link{}
	for _, node := range linkNodes(doc) {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = Text(n)
	return ret
}

func Text(n *html.Node) string {
	if n.Type == html.TextNode {

		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += Text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
