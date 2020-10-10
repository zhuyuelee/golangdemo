package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe("localhost:8080", nil)

}

func root(w http.ResponseWriter, r *http.Request) {
	var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
	<html><body>
	<h1>URL SHORTENER</h1>
	{{if .}}{{.}}<br /><br />{{end}}
	<form action="/short" type="POST">
	Shorten this: <input type="text" name="longUrl" />
	<input type="submit" value="Give me the short URL" />
	</form>
	<br />
	<form action="/long" type="POST">
	Expand this: http://goo.gl/<input type="text" name="shortUrl" />
	<input type="submit" value="Give me the long URL" />
	</form>
	</body></html>
	`))

	rootHtmlTmpl.Execute(w, nil)
}
