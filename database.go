package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func adduser(Rno string, name string) {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS register (RNo TEXT, Name TEXT, date TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO register (Rno, Name, date) VALUES (?, ?, ?)")
	var time string
	statement.Exec(Rno, name, time)

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
