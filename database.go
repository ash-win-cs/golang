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
	rows, _ := database.Query("SELECT datetime('now') as 'Current date & time';")
	var time string
	rows.Scan(&time)
	statement, _ = database.Prepare("INSERT INTO register (Rno, Name, date) VALUES (?, ?, ?)")
	statement.Exec(Rno, name, time)

}

func checkfield(number string, uname string) int {
	var status int = 0
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	rows, _ := database.Query("SELECT Name FROM database WHERE Rno=\"%s\"", number)
	var Name string
	rows.Scan(&Name)
	if uname == Name {
		adduser(number, Name)
		status = 1
	}

	return (status)
}

func main() {

	if checkfield("NSS17EC036", "Ashwin C S") == 0 {
		fmt.Println("No Users Found!!")
	} else {
		fmt.Println("welcome")
	}

}
