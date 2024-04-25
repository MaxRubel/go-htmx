package routes

import (
	"net/http"
	"text/template"
)

func CreateNewWorkout(writer http.ResponseWriter, response *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/template/createNewWorkout.html"))
	tmpl.Execute(writer, nil)
}
