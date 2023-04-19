package bapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/gorilla/mux"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

//tested
func (a *App) AddEventHandler(w http.ResponseWriter, r *http.Request) {
	var newEvent models.Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	/* Need to add an exists function to check if the event already exists
	 * Before implementing the exists need get
	 * Get will check based on event name
	 */
	exists := a.EventExists(newEvent.Event) //, w, r)
	if exists {
		http.Error(w, "Event already exists", http.StatusBadRequest)
		return
	}
	//Addd db helper of added events to Edb
	err = a.CreateEvent(&newEvent, w, r)
	if err != nil {
		fmt.Println("Event failed to be added")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEvent)
	return
}

