package main

import (
	"flag"
	"fmt"
	"gocourse/htmllink"
	"net/http"
)

var exapleHtml = `<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	// reader := strings.NewReader(exapleHtml)
	url := flag.String("url", "https://www.sreyas.com/", "Enter site url")
	flag.Parse()
	links := parseLinks(*url)
	fmt.Printf("%+v\n", links)
}

func parseLinks(url string) []htmllink.Link {

	resp, _ := http.Get(url)

	links, err := htmllink.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return links
}
