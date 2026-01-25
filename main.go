package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.ListenAndServe(":9090", nil)
}
