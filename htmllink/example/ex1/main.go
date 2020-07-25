package main

import (
	"fmt"
	"gocourse/htmllink"
	"strings"
)

var exapleHtml = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page 
  <span> some span </span> </a>
  <a href="/page-two">A link to second page</a>
</body>
</html>
`

func main() {
	reader := strings.NewReader(exapleHtml)
	links, err := htmllink.Parse(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
