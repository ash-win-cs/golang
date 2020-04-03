package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (adno INTEGER PRIMARY KEY, firstname TEXT, semester TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (adno, firstname, semester) VALUES (?, ?, ?)")
	statement.Exec(170120, "ashwin", "s6")
	rows, _ := database.Query("SELECT adno, firstname, semester FROM people")
	var id int
	var firstname string
	var semester string
	for rows.Next() {
		rows.Scan(&id, &firstname, &semester)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
