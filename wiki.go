package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("views/"+ tmpl + ".html")
	t.Execute(w, p)
}
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is Nick's sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

