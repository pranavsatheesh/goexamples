package main

import (
	"flag"
	"fmt"
	"gocourse/books"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var defaultHandlerTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

func main() {
	filename := flag.String("file", "gopher.json", "The Josn file with Stories")
	port := flag.Int("port", 8080, "The port to start the server")
	flag.Parse()
	fmt.Printf("Using the story in %s \n", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := books.JsonStory(f)
	if err != nil {
		panic(err)
	}
	tpl := template.Must(template.New("").Parse(defaultHandlerTemplate))
	h := books.NewHandler(story,
		books.WithTemplate(tpl),
		books.WithPathFn(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the serever on port : %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h), mux)
}
func pathFn(r *http.Request) string {
	fmt.Println(r.URL.Path)
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}
