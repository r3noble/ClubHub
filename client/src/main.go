package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/initializers"
)

/*
	type User struct {
		gorm.Model
		username string `json:"username" gorm:"primary_key"`
		name     string `json:"name"`
		pass     string `json:"pass"`
	}

	func main() {
		port := ":8080"
		router := mux.NewRouter()

		router.HandleFunc("/signin", Signin).Methods("PUT")
		router.HandleFunc("/signup", Signup).Methods("POST")
		http.HandleFunc("/signin", Signin)
		http.HandleFunc("/signup", Signup)
		//initialize databse
		//initDB()
		//start server
		log.Fatal(http.ListenAndServe(port, nil))
	}
*/

var (
	router *mux.Router
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)

	}

	initializers.ConnectDB(&config)

	router = mux.NewRouter()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router.HandleFunc("/api/login",
		func(w http.ResponseWriter, r *http.Request) {
			message := "Welcome to ClubHub login using Golang with Gorm and Postgres"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status": "success", "message": "` + message + `"}`))
		}).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+config.ServerPort, router))

}
