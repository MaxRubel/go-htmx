package queries

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func TestDb() {
	database, _ := sql.Open("sqlite", "./db.sqlite3")
	rows, err := database.Query("SELECT test FROM test")

	if err != nil {
		panic(err)
	}

	defer database.Close()
	defer rows.Close()

	var test string

	for rows.Next() {
		rows.Scan(&test)
		fmt.Println(test)
	}

}
