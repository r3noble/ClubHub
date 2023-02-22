package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/initializers"
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
/*
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
*/
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Age      int
}

//var userMap map[int]User

func main() {
	//	userMap := make(map[int]User)

	var cole User

	cole.ID = 1
	cole.Name = "cole"
	cole.Age = 21
	cole.Email = "cole@rottenberg.org"
	cole.Password = "pass"
	//userMap[1] = cole

	router := mux.NewRouter()
	/*config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	*/
	/*
		//**figure out what handler to send to
		router.HandleFunc("/api/login",
			func(w http.ResponseWriter, r *http.Request) {
				message := "Welcome to ClubHub login using Golang with Gorm and Postgres"
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"status": "success", "message": "` + message + `"}`))
			}).Methods("GET")

		log.Fatal(http.ListenAndServe(":"+config.ServerPort, router))
	*/
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/getSlice", testJSON).Methods("GET")
	router.HandleFunc("/getMap1", func(w http.ResponseWriter, r *http.Request) {
		testJSONMap(w, r, cole)
	}).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

type Response struct {
	Users []User `json:"users"`
}

func testMap() []User {
	//no matter what ds we put users in we will always send user json values as a slice of users
	var users []User
	var user User
	users = append(users, user)
	return users
}

func testUser(user User) []User {
	var users []User
	users = append(users, user)

	return users
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//next function writes back to the response
	fmt.Fprintf(w, "API is running")
}
func testJSON(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var response Response
	//setting the content type to json
	w.Header().Set("Content-Type", "application/json")
	users := testMap()
	response.Users = users

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
func testJSONMap(w http.ResponseWriter, r *http.Request, user User) {
	w.WriteHeader(http.StatusOK)
	var response Response
	//setting the content type to json
	w.Header().Set("Content-Type", "application/json")
	users := testUser(user)
	response.Users = users

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
