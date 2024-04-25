package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"testing.com/queries"
)

type Film struct {
	Title string
	Director string
}

func main() {
	fmt.Println("Running Server on port 42069")
	
	//handler funcs
	homepage := func (writer http.ResponseWriter, response *http.Request){
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(writer, films)
	}

	postNewFilm := func (writer http.ResponseWriter, response *http.Request){
		title:= response.PostFormValue("title")
		director := response.PostFormValue("director")
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(writer, "film-list-element", Film{Title: title, Director: director})
	}

	//request handlers
	http.HandleFunc("/", homepage)
	http.HandleFunc("/add-film/", postNewFilm)

	queries.GetAllPeople()


	log.Fatal(http.ListenAndServe(":42069", nil))
}