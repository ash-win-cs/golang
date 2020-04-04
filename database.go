package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func adduser() {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS register (RNo TEXT, Name TEXT, date TEXT)")
	statement.Exec()

}

func checkfield(number string) {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	rows, _ := database.Query("SELECT Name FROM database WHERE Rno=\"%s\"", number)
	var Name string
	rows.Scan(&Name)
	fmt.Println(Name)

}

func main() {

}
