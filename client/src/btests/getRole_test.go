package main

import (
	"testing"
    "net/http"
    "net/http/httptest"
	"fmt"
	"bytes"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestGetRoleHandler(t *testing.T) {
    // Setup test data
    user := models.User{
		ID: "111",
        Name: "test_user",
        Clubs: "club1",
    }
    club := models.Club{
        Name: "club1",
    }

	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	testCdb, err := gorm.Open(sqlite.Open("testClub.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

    // Create a new request
    reqBody := fmt.Sprintf(`{"id":"%s","name":"%s"}`, user.Name, club.Name)
    req, err := http.NewRequest("POST", "/api/getRole", bytes.NewBufferString(reqBody))
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
    }

    // Create a new response recorder
    rr := httptest.NewRecorder()

    // Call the handler function
    a := &bapp.App{
        DB:  testDB,
		Cdb: testCdb,
    }
    handler := http.HandlerFunc(a.GetRoleHandler)
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect
    expectedResp := models.ClubAdder{
        Name: "testCreate",
        ID:   "Not a Member",
    }
    expectedJSON, _ := json.Marshal(expectedResp)
    if rr.Body.String() != string(expectedJSON) {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expectedJSON))
    }
}
