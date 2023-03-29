package main

import(
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    
    "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    "encoding/json"
)

/*func TestLoginHandler(t *testing.T) {
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
    
}*/

func TestLoginHandler(t *testing.T) {
    testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    //migrate db schema
    err = testDB.AutoMigrate(&User{})
    if err != nil {
        t.Fatal(err)
    }

    // Create a new App instance with the mock database
    a := &App{
        db: testDB,
        r:  mux.NewRouter(),
    }

    //create a mock user to use for authentication
    user := &User{
        ID: "123",
        Name:     "testuser",
        Email:    "testuser@example.com",
        Password: "testpassword",
    }
    err = a.db.Create(user).Error
    if err != nil {
        t.Fatal(err)
    }

    //create a mock request to server
    body := []byte(`{"username": "testuser", "password": "testpassword"}`)
    req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    //create mock request recorder
    mockRec := httptest.NewRecorder()

    //call loginHandler with mock request recorder
    a.r.HandleFunc("/api/login", a.loginHandler)
    a.r.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code; status != http.StatusOK {
        t.Errorf("Wrong status returned, got %v, want %v", status, http.StatusOK)
    }

    //check response body
    var responseUser User
    err = json.NewDecoder(mockRec.Body).Decode(&responseUser)
    if err != nil {
        t.Fatal(err)
    }
    if responseUser.ID != user.ID || responseUser.Name != user.Name || responseUser.Email != user.Email {
        t.Errorf("Returned unexpected user data, got %v, want %v", responseUser, user)
    }
}

