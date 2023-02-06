package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var vid int
var vname string

func main() {
	//createDB()
	QuerySelect()
}

func QuerySelect() {
	const file string = "activities.db"
	db, _ := sql.Open("sqlite3", file)
	defer db.Close()

	_ = db.QueryRow("SELECT id, description FROM activities WHERE id=1").Scan(&vid, &vname)
	defer db.Close()
	fmt.Printf(vname)
}

func createDB() {
	const create string = `
	CREATE TABLE IF NOT EXISTS activities(
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT
	);`
	const file string = "activities.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		println(err.Error())
	}
	if _, err := db.Exec(create); err != nil {
		println(err.Error())
	}

}
