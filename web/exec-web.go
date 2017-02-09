package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const (
	SUFFIX       = ".txt"
	PERMISSION   = 0600
	LISTEN_PORT  = ":8080"
	TEMPLATE_DIR = "tmpl"
	FILE_DIR     = "data"
	HOMEPAGE     = "frontpage"
)

var (
	templates *template.Template
	validPaths *regexp.Regexp
)

type Page struct {
	Title string
	Body  []byte
}

// Save the body to file
func (p *Page) save() error {
	filename := p.Title + SUFFIX
	permission := os.FileMode(PERMISSION)
	return ioutil.WriteFile(FILE_DIR+"/"+filename, p.Body, permission)
}

// Read body content from file
func loadPage(title string) (*Page, error) {
	filename := FILE_DIR + "/" + title + SUFFIX
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Read the template from the templates cache
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Do some common func in the head of each handlers
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPaths.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if !validPaths.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/view/"+HOMEPAGE, http.StatusFound)
}

func main() {
	// Cache the templates
	templates = template.Must(template.ParseFiles(
		TEMPLATE_DIR+"/edit.html",
		TEMPLATE_DIR+"/view.html",
	))

	validPaths = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]*)$|^/$")

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(LISTEN_PORT, nil)
}
