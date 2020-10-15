package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
var templates = make(map[string]*template.Template)
var err error

//Page 页面
type Page struct {
	Title string
	Body  []byte
}

func init() {
	for _, tmpl := range []string{"edit", "view", "index"} {
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}

func main() {
	http.HandleFunc("/", makeHandler(indexHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	url := "localhost:8080"
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.TrimRight(path, "/")
		if path == "" {
			path = "index"
		}
		title := path[strings.LastIndex(path, "/")+1:]
		fmt.Printf("title=[%s]\n", title)
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("path=", r.URL.Path)
	p, err := load(title)
	if err != nil { // page not found
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "index", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("path=", r.URL.Path)
	p, err := load(title)
	if err != nil { // page not found
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("path=", r.URL.Path)
	p, err := load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("path=", r.URL.Path)
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// file created with read-write permissions for the current user only
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return &Page{Title: title, Body: body}, nil
}
