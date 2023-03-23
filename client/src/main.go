package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/gorilla/mux"
	//"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/initializers"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// var userMap map[int]User
type App struct {
	db *gorm.DB
	u  map[string]User
	r  *mux.Router
	mu sync.Mutex
}

func WriteOnceMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{w, false}
		h.ServeHTTP(wrappedWriter, r)
		if !wrappedWriter.wroteHeader {
			wrappedWriter.WriteHeader(http.StatusOK)
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.wroteHeader = true
}

func (a *App) start() {
	a.r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	a.r.HandleFunc("/api/health", HealthCheck).Methods("GET")
	//query-based matching using id
	a.r.HandleFunc("/api/getUser/{id}", a.IdHandler).Methods("GET")
	a.r.HandleFunc("/api/addUser", a.AddUserHandler).Methods("POST")
	a.r.HandleFunc("/api/login", a.loginHandler).Methods("POST") // handlers login
	http.ListenAndServe(":8080", a.r)
}
func main() {
	//Initialize and open DB here
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Error in opening DB")
	}
	//calls AutoMigrate and throws error if cannot migrate
	//formats db to replicate user struct
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Error in migrating db")
	}
	app := App{
		db: db,
		u: make(map[string]User),
		r: mux.NewRouter(),
	}
	app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

	app.start()
}

func (a *App) CreateUser(user *User, w http.ResponseWriter, r *http.Request) error {
	//err := a.db.Create(user).Error
	err := a.db.Model(&User{}).Create(user).Error
	fmt.Println(user.Name)
	if err != nil {
		fmt.Printf("Error creating user: %s", err.Error())
		http.Error(w, "Could not insert user into database", http.StatusInternalServerError)
		return err
	}
	return nil
}

//called to search for user when adding user, does not return 404 if user not found as this is the desired result
func (a *App) UserExists(name string, w http.ResponseWriter, r *http.Request) *User {
	//call is based on User Strcut not credentials struct, may need to change
	user := User{}
	if err := a.db.First(&user, User{Name: name}).Error; err != nil {
		//respondError(w, http.StatusNotFound, err.Error())
		fmt.Println("User not located, adding to database...")
		return nil
	}
	return &user
}

//searches DB for user, returns nil if none found
func (a *App) QueryDbByID(id string, w http.ResponseWriter, r *http.Request) *User {
	//call is based on User Strcut not credentials struct, may need to change
	user := User{}
	if err := a.db.First(&user, User{ID: id}).Error; err != nil {
		//respondError(w, http.StatusNotFound, err.Error())
		http.Error(w, "User not located", http.StatusNotFound)
		return nil
	}
	 return &user
}

//searches DB fpr user, returns nil if none found
func (a *App) QueryByName(name string, w http.ResponseWriter, r *http.Request) *User {
	//call is based on User Strcut not credentials struct, may need to change
	user := User{}
	if err := a.db.First(&user, User{Name: name}).Error; err != nil {
		//respondError(w, http.StatusNotFound, err.Error())
		fmt.Printf("Error: %s", err.Error())
		//http.Error(w, "User not located", http.StatusNotFound)
		return nil
	}
	return &user
}

func (a *App) GetUserByName(name string, w http.ResponseWriter, r *http.Request) (*User, error){
	fmt.Println("Entering GetUserByName")
	user := a.QueryByName( name, w, r)
	if user == nil {
		return nil, fmt.Errorf("user with name %d not found", name)
	}
	return user, nil
}

func (a *App) GetUserByID(id string, w http.ResponseWriter, r *http.Request) (*User, error) {
	fmt.Println("Entering GetUserByID")
	//TREY: QUERY DB HERE FOR USER ID (Call QueryDbByID)
	user := a.QueryDbByID(id, w, r)
	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST and the URL path is /user/login
	// Decode the JSON payload from the request body
	fmt.Println("Successfully entered Login Handler")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Bad Json in Body")
		return
	}
	// print the request body
    fmt.Printf("Request body: %+v\n", creds)

	// Check if the required fields (username and password) are present
	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Authenticate the user using the provided credentials (not shown)
	// ...
	//TREY: QUERY DB here for username
	user := a.QueryByName(creds.Username, w, r)
	if user == nil {
		http.Error(w, "Invalid Username", http.StatusUnauthorized)
		fmt.Println("No found user")
		return
	}
	//now we check the password
	knownPass := user.Password
	if knownPass != creds.Password {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		fmt.Println("No found password")
		return
	}
	/*response := struct {
		Message string `json:"message"`
	}{
		Message: "Login successful",
	}*/

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("About to pass back user")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	fmt.Println("Passing back success")
	// Send a success response
	return

	// Send a 404 Not Found response if the URL path doesn't match
}
func (a *App) IdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	id := vars["id"]
	// Look up the user with the given id in the map
	//TREY: Get user by ID must be updated for DB support
	user, err := a.GetUserByID(id, w, r)
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

	// print the request body
    fmt.Printf("Request body: %+v\n", newUser)

	newUser.ID = strconv.Itoa(rand.Intn(1000))

	// Check if the user ID already exists in the map
	//TREY: Query DB for ID, if EXISTS, print same error
	if user := a.UserExists(newUser.ID, w, r); user != nil {
		http.Error(w, "User with that ID already exists", http.StatusBadRequest)
		return
	}

	// Add the new user to the map
	//TREY: Call function to add new user to db
	err = a.CreateUser(&newUser, w, r)
	if err != nil {
		fmt.Println("User Unsuccessfully added to DB")
	}
	else{
		fmt.Printf("User successfully created with name %s and ID %s", newUser.Name, newUser.ID)
	}

	// Return the new user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//next function writes back to the response
	fmt.Println("Health check accessed")
	fmt.Fprintf(w, "API is running")
}

func (a *App) profileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the username parameter from the URL path
	username := r.URL.Query().Get("username")

	// Retrieve the profile data from the map
	//TREY: QUERY DB for username
	profile, _ := a.GetUserByName(username, w, r)
	if profile == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the profile data to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
