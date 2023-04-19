package bapp

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)
func (a *App) GetRoleHandler(w http.ResponseWriter, r *http.Request){
	//get user by json in POST
	vars := mux.Vars(r)
	name := vars["id"]
	clubName := vars["name"]
	var user models.User
	uerr := a.DB.First(&user, models.User{Name: name})
	if uerr != nil {
		http.Error(w, "User DNE", http.StatusBadRequest)
	}
	var club models.Club
	cerr := a.Cdb.First(&club, models.Club{Name: clubName})
	if cerr != nil {
		http.Error(w, "Club DNE", http.StatusBadRequest)
	}
	//Add user name to the ClubAdder as response
	var resp models.ClubAdder
	resp.Name = user.Name
	var inv string
	list := user.Clubs
	if list == "" {
		inv = "Not a member"
	}
	var tmp string
	for i := 0; i < len(user.Clubs); i++ {
		if string(user.Clubs[i]) == "," {
			tmp = ""
		}
		tmp += string(user.Clubs[i])
		//if yes -> send back message to frontend to display error message
		if tmp == clubName {
			resp.ID = "Member"
			jsonResponse, founderr := json.Marshal(resp)
			if founderr != nil {
				http.Error(w, "Issue Marshaling found club", http.StatusInternalServerError)
				fmt.Println("Issue marshaling ClubAdder as Response")
				return
			}

			//send back success message to front end
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
			return
		}
	}

	resp.ID = inv
	jsonResponse, finerr := json.Marshal(resp)
	if finerr != nil {
		http.Error(w, finerr.Error(), http.StatusInternalServerError)
		fmt.Println("Issue marshaling ClubAdder as Response")
		return
	}

	//send back success message to front end
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	
	//list will hold list of User clubs
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

	var clubList string
	//create new string of clubs to be added
	if user.Clubs == "" {
		clubList = ident.Name
	} else {
		clubList = user.Clubs + "," + ident.Name
	}
	//determine if they are in the club already
	var clubName string
	for i := 0; i < len(user.Clubs); i++ {
		if string(user.Clubs[i]) == "," {
			clubName = ""
		}
		clubName += string(user.Clubs[i])
		//if yes -> send back message to frontend to display error message
		if clubName == ident.Name {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("User already a part of that club")
			return
		}
	}
	//otherwise-> edit the club column for their userDB slot
	user.Clubs = clubList
	a.DB.Model(&models.User{}).Where("ID=?", user.ID).Update("Clubs", clubList)

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

//tested
func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST and the URL path is /user/login
	// Decode the JSON payload from the request body
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Successfully entered Login Handler")
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Bad Json in Body")
		return
	}

	// Check if the required fields (username and password) are present
	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Authenticate the user using the provided credentials (not shown)
	// ...
	//TREY: QUERY DB here for username
	user := a.QueryByName(creds.Email, w, r)
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

//tested
func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new user data
	var newUser models.User
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
	fmt.Printf("User successfully created with name %s and ID %s\n", newUser.Name, newUser.ID)

	// Return the new user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// NOT CURRENTLY USED jenna is trying to adapt this
func (a *App) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the username parameter from the URL path
	vars := mux.Vars(r)
	name := vars["id"]

	// Retrieve the profile data from the map
	//TREY: QUERY DB for username
	profile, _ := a.GetUserByNName(name, w, r)
	if profile == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Convert the profile data to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//next function writes back to the response
	fmt.Println("Health check accessed")
	fmt.Fprintf(w, "API is running")
}
