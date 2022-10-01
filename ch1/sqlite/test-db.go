package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// go get -u "github.com/mattn/go-sqlite3"
)

func main() {
	println("Starting sqlite3 memory db")
	db, err := sql.Open("sqlite3", ":memory:")
	if nil != err {
		panic(err)
	}
	defer db.Close()

	queries := []string{
		"select CURRENT_TIMESTAMP", 
		"select date()","select time()"}

	var result string
	for _, qry := range queries {
		println(qry)
		row := db.QueryRow(qry)
		err = row.Scan(&result)
		if nil != err {
			println("Error found: " + err.Error())
			continue
		}
		println(result)
	}
	println("Completed sqlite3 memory db")

}
