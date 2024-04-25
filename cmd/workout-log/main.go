package main

import (
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"testing.com/internal/queries"
	"testing.com/internal/routes"
)

func main() {
	fmt.Println("Running Server on port 42069...")
	queries.TestDb()

	//request handlers
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/create-new-workout/", routes.CreateNewWorkout)
	http.HandleFunc("/post-new-workout/", queries.PostNewWorkOut)
	http.HandleFunc("/view-all-workouts/", routes.ViewAllWorkouts)
	http.HandleFunc("/view-single-workout/", routes.ViewAllWorkouts)

	log.Fatal(http.ListenAndServe(":42069", nil))
}
