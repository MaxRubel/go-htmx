package routes

import (
	"net/http"
	"text/template"

	"testing.com/internal/queries"
)

func ViewAllWorkouts(writer http.ResponseWriter, response *http.Request) {
	Workouts := queries.GetAllWorkouts()
	tmpl := template.Must(template.ParseFiles("internal/template/viewAllWorkouts.html"))
	tmpl.Execute(writer, Workouts)
}
