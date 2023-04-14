package bapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

//NOT tested, should have identical functionality to AddUserHandler, minus generation of a random ID
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
	fmt.Printf("Club successfully created with name %s", newClub.Name)

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

func (a *App) JoinClubHandler(w http.ResponseWriter, r *http.Request) {
	//get identification of user to be accessed
	var ident models.ClubAdder
	err := json.NewDecoder(r.Body).Decode(&ident)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//get user to be accessed from DB
	var user models.User
	if err = a.DB.First(&user, "ID=?", ident.ID).Error; err != nil {
		fmt.Println("User with that name not found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//create new string of clubs to be added
	clublist := user.Clubs + "," + ident.Name
	//determine if they are in the club already
	var clubName string
	for i := 0; i < len(user.Clubs); i++ {
		if string(user.Clubs[i]) == ","{
			clubName = ""
		}
		clubName += string(user.Clubs[i])
		//if yes -> send back message to frontend to display error message
		if clubName == ident.Name{
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("User already a part of that club")
			return
		}
	}
	//otherwise-> edit the club column for their userDB slot
	user.Clubs = clublist
	a.DB.Model(&models.User{}).Where("ID=?", user.ID).Update("Clubs", clublist)
	
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	//send back success message to front end
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
