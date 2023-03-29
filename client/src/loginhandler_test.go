package main

import(
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestLoginHandler(t *testing.T) {
    a := &App{
        u: map[string]User{
            "1": {ID: 1, Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"},
        },
    }
    //create mock request
    body := []byte(`{"username": "testuser", "password": "testpass"}`)
    req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    //create mock response recorder
    mockRec := httptest.NewRecorder()

    //call loginHandler with mock request recorder
    http.HandleFunc("/api/login", a.loginHandler)
    http.DefaultServeMux.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code ;status != http.StatusOK {
        t.Errorf("wrong status returned, got %v, want %v", status, http.StatusOK)
    }

    //now check response body
    expected := `{"message": "Login successful"}`
    if mockRec.Body.String() != expected {
        t.Errorf("returned unexpected body, got %v want %v", mockRec.Body.String(), expected)
    }
    
}

func TestAddUserHandler(t *testing.T) {
    testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }
    defer testDB.close()

    //migrate db schema
    err = testDB.AutoMigrate(&User{})
    if err != nil {
        t.Fatal(err)
    }

    // Create a new App instance with the mock database
	a := &App{
		db: db,
		r:  mux.NewRouter(),
	}

    //create a mock request to server
    body:= []byte(`{"ID": "69", "Name": "tester", "Email": "fml@fm.com", "Password": "testP"}`)
    req, err := http.NewRequest("POST", "/api/addUser", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    //create mock request recorder
    mockRec := httptest.NewRecorder()

    //call addUser handler with mock request recorder
    http.HandleFunc("api/addUser", a.AddUserHandler)
    http.DefaultServeMux.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code; status != http.StatusOK {
        t.Errorf("Wrong status returned, got %v, want %v", status, http.StatusOK)
    }

    expected := `{"message": "Login successful"}`
    if mockRec.Body.String() != expected {
        t.Errorf("Returned unexpected body, got %v, want %v", mockRec.Body.String(), expected)
    }
}

