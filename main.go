package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title string
	Director string
}

func main() {
	
	port := 8000
    fmt.Printf("Server running on http://localhost:%d",port)

	home := func (w http.ResponseWriter, r *http.Request)  {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films":{
				{Title: "John Wick",Director: "Joko Anwar"},
				{Title: "John Wick 4",Director: "Ustad Mansur"},
				{Title: "The Thing",Director: "John Carpenter"},
			},
		}

		tmpl.Execute(w, films)
	}

	add_film := func (w http.ResponseWriter, r *http.Request)  {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>",title, director)
		tmpl,_ := template.New("film").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/",home)
	http.HandleFunc("/add-film/",add_film)


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port),nil))
}
