package main

import (
	"fmt"
	"gocourse/htmllink"
	"strings"
)

var exapleHtml = `<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
