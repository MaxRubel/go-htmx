package routes

import (
	"net/http"
	"text/template"
)

func HomePage(writer http.ResponseWriter, response *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/template/index.html"))
	tmpl.Execute(writer, nil)
}
