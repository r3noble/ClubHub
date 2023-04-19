package main

import (
	"testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    "encoding/json"

    "github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestAddEventHandler(t *testing.T) {
    testEdb, err := gorm.Open(sqlite.Open("testEvent.db"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    //migrate db schema
    err = testEdb.AutoMigrate(&models.Event{})
    if err != nil {
        t.Fatal(err)
    }

    // Create a new App instance with the mock database
    a := &bapp.App{
        Edb: testEdb,
        R:  mux.NewRouter(),
    }
     //create a mock user to use for authentication
    event := models.Event{
    	Club: "testEvent",
        Event: "T1",
        Date: "T2",
		StartTime: "T3",
        EndTime: "T4",
		Users: "many",
    }
    body, _ := json.Marshal(event)
    req, err := http.NewRequest("POST", "/api/addEvent", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    //mock request recorder
    mockRec := httptest.NewRecorder()

    //make request with mockRec and response
    a.R.HandleFunc("/api/addEvent", a.AddEventHandler)
    a.R.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code; status != http.StatusOK {
        t.Errorf("Wrong status returned, got %v want %v", status, http.StatusOK)
    }
    //check response body
    var responseEvent models.Event
    err = json.NewDecoder(mockRec.Body).Decode(&responseEvent)
    if err != nil {
        t.Fatal(err)
    }

    if responseEvent.Club != event.Club || responseEvent.Event != event.Event || responseEvent.Date != event.Date || responseEvent.StartTime != event.StartTime || responseEvent.EndTime != event.EndTime {
        t.Errorf("Returned unexpected user data, got %v, want %v", responseEvent, event)
    }


}