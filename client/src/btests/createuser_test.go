package main

import(
    "testing"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestCreateUser(t *testing.T){
	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//migrate db schema
	err = testDB.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatal(err)
	}

	//create app instance w/ mock db
	a := &bapp.App{
		DB: testDB,
		R: mux.NewRouter(),
	}

	//create mock user to add
	user := &models.User {
		ID: "001",
		Name: "testCreate",
		Email: "testlogin@test.com",
		Password: "tlpword",
	}

	mockW := httptest.NewRecorder()
    mockR := httptest.NewRequest(http.MethodPost, "/users", nil)

	//call create user on testDB
	err = a.CreateUser(user, mockW, mockR)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	//check user was added to db
	var created models.User
	a.DB.First(&created, user.ID)
	if created != *user{
		t.Errorf("Expected user %+v, but got %+v", user, &created)
	}

	//check http status code
	if mockW.Code != http.StatusOK{
		t.Errorf("Expected status code %d but got %d", http.StatusOK, mockW.Code)
	}
}