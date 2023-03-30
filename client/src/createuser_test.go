package main

import(
    "testing"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T){
	testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//migrate db schema
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}

	//create app instance w/ mock db
	a := &App{
		db: testDB,
		r: mux.NewRouter(),
	}

	//create mock user to add
	user := &User {
		ID: "111",
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
	var created User
	a.db.First(&created, user.ID)
	if created != *user{
		t.Errorf("Expected user %+v, but got %+v", user, &created)
	}

	//check http status code
	if mockW.Code != http.StatusOK{
		t.Errorf("Expected status code %d but got %d", http.StatusOK, mockW.Code)
	}
}