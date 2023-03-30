package main

import(
    "testing"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func TestUserExists(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//migrate DB schema
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}

	//create app instance w/ mock db
	a := &App {
		db: testDB,
		r: mux.NewRouter(),
	}

	//create user to add to db, this will be checked for later
	user := &User {
		ID: "01",
		Name: "testExists",
		Email: "testExists@test.com",
		Password: "pass123",
	}

	err = a.db.Create(user).Error
	if err != nil {
		t.Errorf("Error in adding, expected none, got %s", err)
	}

	mockW := httptest.NewRecorder()
	mockR := httptest.NewRequest(http.MethodPost, "/users", nil)

	exists := a.UserExists(user.Name, mockW, mockR)
	if !exists {
		t.Errorf("Expected user %s to exists but it does not", user.Name)
	}
}