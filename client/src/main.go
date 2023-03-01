package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

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
	u  map[int]User
	r  *mux.Router
	mu sync.Mutex
}

func (a *App) start() {
	// ADD DATABASE MIGRATION TO APP instance e.g. a.db.AutoMigrate....
	a.r.HandleFunc("/health", HealthCheck).Methods("GET")
	//query-based matching using id
	a.r.HandleFunc("/user/{id}", a.IdHandler).Methods("GET")
	a.r.HandleFunc("/user/add", a.AddUserHandler).Methods("POST")
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

/*
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
*/
func (a *App) GetUserByID(id int) (*User, error) {
	user, ok := a.u[id]
	if !ok {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return &user, nil
}

func (a *App) IdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	// Look up the user with the given id in the map
	user, err := a.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new user data
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user ID already exists in the map
	if _, ok := a.u[newUser.ID]; ok {
		http.Error(w, "User with that ID already exists", http.StatusBadRequest)
		return
	}

	// Add the new user to the map
	a.mu.Lock()
	defer a.mu.Unlock()
	a.u[newUser.ID] = newUser

	// Return the new user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

type Response struct {
	Users []User `json:"users"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//next function writes back to the response
	fmt.Fprintf(w, "API is running")
}
