package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Println("Error while connect to DB: ", err)
	}

	h, err := Initialize(db)
	if err != nil {
		log.Fatal("Initialize got failure!", err)
	}

	// Sử dụng đối tượng `handler`
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
