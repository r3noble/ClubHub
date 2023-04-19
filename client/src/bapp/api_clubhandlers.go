package bapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

//should have identical functionality to AddUserHandler, minus generation of a random ID
func (a *App) AddClubHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new Club data
	var newClub models.Club
	err := json.NewDecoder(r.Body).Decode(&newClub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//newClub.ID = strconv.Itoa(rand.Intn(1000))

	// Check if the Club ID already exists in the map
	//TREY: Query DB for ID, if EXISTS, print same error
	exists := a.ClubExists(newClub.Name, w, r)
	if exists {
		http.Error(w, "Club with that name already exists", http.StatusBadRequest)
		return
	}

	// Add the new Club to the map
	//TREY: Call function to add new Club to db
	err = a.CreateClub(&newClub, w, r)
	if err != nil {
		fmt.Println("Club Unsuccessfully added to DB")
	}
	fmt.Printf("Club successfully created with name %s\n", newClub.Name)

	// Return the new Club data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newClub)
}

func (a *App) GetClubHandler(w http.ResponseWriter, r *http.Request){
	//Get all clubs from the clubs database to check if exist
	//var tmp models.Club
	vars := mux.Vars(r)
	name := vars["id"]

	club := models.Club{}
	if err := a.Cdb.First(&club, models.Club{Name: name}).Error; err != nil {
		fmt.Println("Club not located, adding to database...")
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(club)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	return
}

