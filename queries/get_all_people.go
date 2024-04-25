package queries

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

func GetAllPeople(){
	database, _ := sql.Open("sqlite", "./test.sqlite3")
	rows, err := database.Query("SELECT id, firstname, lastname FROM people")
	
	if err !=nil{
		panic(err)
	}

	defer database.Close()
	defer rows.Close()

	var id int
	var firstname string
	var lastname string
	
	for rows.Next(){
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id), firstname + " " + lastname)
	}


}