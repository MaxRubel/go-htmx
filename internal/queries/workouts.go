package queries

import (
	"database/sql"
	"net/http"

	structs "testing.com/internal/models"
)

func PostNewWorkOut(writer http.ResponseWriter, response *http.Request) {

	name := response.PostFormValue("name")
	category := response.PostFormValue("category")

	database, err := sql.Open("sqlite", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	stmt, err := database.Prepare("INSERT INTO workouts (name, category) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}

	defer database.Close()
	defer stmt.Close()

	_, err = stmt.Exec(name, category)
	if err != nil {
		panic(err)
	}
}

func GetAllWorkouts() []structs.Workout {

	database, err := sql.Open("sqlite", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	rows, err := database.Query("SELECT id, name, category FROM workouts")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	defer rows.Close()

	var workouts []structs.Workout

	for rows.Next() {
		var name, category string
		var id int
		err := rows.Scan(&id, &name, &category)
		if err != nil {
			panic(err)
		}
		workouts = append(workouts, structs.Workout{Id: id, Name: name, Category: category})
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
	return (workouts)
}
