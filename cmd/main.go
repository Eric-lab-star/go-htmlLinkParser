package main

import (
	"fmt"
	"strings"

	htmlLinks "github.com/Eirc-lab-star/htmlLinkParser/pkg/htmlLinks"
)

var helloHTML string = `
	<html>
		<body>
			<h1>Hello!</h1>
			<a href="/other-page">A link to another page</a>
		</body>
	</html>
`

func main() {
	io := strings.NewReader(helloHTML)
	links, err := htmlLinks.Parse(io)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)

}
