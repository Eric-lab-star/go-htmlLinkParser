package htmlLinks

import "io"

// Link represent a link in an HTML document.
type Link struct {
	Text string
	Href string
}

// Parse wil take in an HTML documment and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]string, error) {
	return nil, nil
}
