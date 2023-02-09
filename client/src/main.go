package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "kraus"
	password = "0228"
	dbname   = "mydb"
)

var db *sql.DB

func main() {
	port := ":8080"
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	//initialize databse
	//initDB()
	//start server
	log.Fatal(http.ListenAndServe(port, nil))
}

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("kraus", psqlInfo)
	if err != nil {
		fmt.Println("Error opening connection to database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
}
