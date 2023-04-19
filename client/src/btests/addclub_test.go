package main

import (
	"testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    "fmt"
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    "encoding/json"

    "github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestAddClubHandler(t *testing.T) {
    testCdb, err := gorm.Open(sqlite.Open("testClub.db"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    //migrate db schema
    err = testCdb.AutoMigrate(&models.Club{})
    if err != nil {
        t.Fatal(err)
    }

    // Create a new App instance with the mock database
    a := &bapp.App{
        Cdb: testCdb,
        R:  mux.NewRouter(),
    }
     //create a mock user to use for authentication
    club := models.Club{
    	Name: "testclub",
        President: "T1",
        VP: "T2",
		Treasurer: "T3",
        About: "hey",
    }
    body, _ := json.Marshal(club)
    req, err := http.NewRequest("POST", "/api/addClub", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }
    //mock request recorder
    mockRec := httptest.NewRecorder()

    //make request with mockRec and response
    a.R.HandleFunc("/api/addClub", a.AddClubHandler)
    a.R.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code; status != http.StatusOK {
        t.Errorf("Wrong status returned, got %v want %v", status, http.StatusOK)
    }
    fmt.Println("API Request made")
    //check response body
    var responseClub models.Club
    err = json.NewDecoder(mockRec.Body).Decode(&responseClub)
    if err != nil {
        t.Fatal(err)
    }

    if responseClub.Name != club.Name || responseClub.President != club.President || responseClub.VP != club.VP || responseClub.Treasurer != club.Treasurer || responseClub.About != club.About {
        t.Errorf("Returned unexpected user data, got %v, want %v", responseClub, club)
    }
}