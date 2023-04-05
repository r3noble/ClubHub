package main

import (
	"testing"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestGetByID(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//migrate DB schema
	err = testDB.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatal(err)
	}

	//create app instance w/ mock db
	a := &bapp.App {
		DB: testDB,
		R: mux.NewRouter(),
	}

	//create user to add to db, this will be checked for later
	user := &models.User {
		ID: "101",
		Name: "testExists",
		Email: "testExists@test.com",
		Password: "pass123",
	}

	err = a.DB.Create(user).Error
	if err != nil {
		t.Errorf("Error in adding, expected none, got %s", err)
	}

	mockW := httptest.NewRecorder()
	mockR := httptest.NewRequest(http.MethodPost, "/users", nil)

	found, err := a.GetUserByID(user.ID, mockW, mockR)

	if found.ID != "101" {
		t.Errorf("Unexpected ID returned, got %s, want 101", found.ID)
	}
}