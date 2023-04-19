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

func TestJoinClub(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//migrating Edb schema
	err = testDB.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatal(err)
	}

	//create app instance
	a := &bapp.App {
		DB: testDB,
		R: mux.NewRouter(),
	}

	request := models.ClubAdder {
		ID: "101",
		Name: "DEEZ",
	}

	body, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", "/api/joinClub", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	mockRec := httptest.NewRecorder()

	a.R.HandleFunc("/api/joinClub", a.JoinClubHandler)
	a.R.ServeHTTP(mockRec, req)

	if status := mockRec.Code; status != http.StatusOK {
		t.Errorf("Wrong status returned, got %v wanted %v", status, http.StatusOK)
	}

	//check is response body matches request
	var responseUser models.User
	err = json.NewDecoder(mockRec.Body).Decode(&responseUser)
	if err != nil {
		t.Fatal(err)
	}

	if responseUser.ID != request.ID || responseUser.Clubs != request.Name {
		t.Errorf("returned unexpected user data, got %v want %v", responseUser, request)
 	}	
}