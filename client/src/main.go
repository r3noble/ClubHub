package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

// var userMap map[int]User
type App struct {
	//db *gorm.DB
	u map[int]User
	r *mux.Router
}

func (a *App) start() {
	// ADD DATABASE MIGRATION TO APP instance e.g. a.db.AutoMigrate....
	a.r.HandleFunc("/health", HealthCheck).Methods("GET")
	http.ListenAndServe(":8080", a.r)
}
func main() {
	//	userMap := make(map[int]User)

	var cole User

	cole.ID = 1
	cole.Name = "cole"
	cole.Age = 21
	cole.Email = "cole@rottenberg.org"
	cole.Password = "pass"
	//userMap[1] = cole

	app := App{
		//db: db,
		u: make(map[int]User),
		r: mux.NewRouter(),
	}
	app.u[1] = User{ID: 1, Name: "Cole", Age: 21, Email: "cole@rottenberg.org", Password: "pass"}
	app.start()

	//router.HandleFunc("/health", HealthCheck).Methods("GET")
	//router.HandleFunc("/getSlice", testJSON).Methods("GET")
	/*router.HandleFunc("/getMap1", func(w http.ResponseWriter, r *http.Request) {
		testJSONMap(w, r, cole)
	}).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
	*/
}
func (a *App) addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var us User
	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	s.ID = uuid.New().String()
	err = a.db.Save(&s).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
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
