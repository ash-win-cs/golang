package main

import (
    "github.com/Equanox/gotron"

    "database/sql"
    "fmt"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
)

func database() {
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

func main() {
    // Create a new browser window instance
    window, err := gotron.New("webapp")
    if err != nil {
        panic(err)
    }

    // Alter default window size and window title.
    window.WindowOptions.Width = 1200
    window.WindowOptions.Height = 980
    window.WindowOptions.Title = "Sandeep K"

    // Start the browser window.
    // This will establish a golang <=> nodejs bridge using websockets,
    // to control ElectronBrowserWindow with our window object.
    done, err := window.Start()
    if err != nil {
        panic(err)
    }
    
    // Open dev tools must be used after window.Start 
    // window.OpenDevTools()
    
    // Wait for the application to close
    <-done
}


