package backend_app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

//NOT CURRENTLY USED, NEEDS UPDATING
func (a *App) AddClubHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new user data
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = strconv.Itoa(rand.Intn(1000))

	// Check if the user ID already exists in the map
	//TREY: Query DB for ID, if EXISTS, print same error
	exists := a.UserExists(newUser.ID, w, r)
	if exists {
		http.Error(w, "User with that ID already exists", http.StatusBadRequest)
		return
	}

	// Add the new user to the map
	//TREY: Call function to add new user to db
	err = a.CreateUser(&newUser, w, r)
	if err != nil {
		fmt.Println("User Unsuccessfully added to DB")
	}
	fmt.Printf("User successfully created with name %s and ID %s", newUser.Name, newUser.ID)

	// Return the new user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}