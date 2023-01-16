package main

import (
	"fmt"
	"strings"

	"github.com/Eirc-lab-star/htmlLinkParser/pkg/htmlLinks"
)

var helloHTML string = `
	<html>
		<body>
			<h1>Hello!</h1>
			<a href="/first-page">A link to another page</a>
			<span>
				<a href="/second-page">A link to another page
					<div>heelo</div>
				</a>
			</span>
			
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
