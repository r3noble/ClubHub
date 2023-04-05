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

    "github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)

func TestAddUserHandler(t *testing.T) {
    testDB, err := gorm.Open(sqlite.Open("testUser.db"), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    //migrate db schema
    err = testDB.AutoMigrate(&models.User{})
    if err != nil {
        t.Fatal(err)
    }

    // Create a new App instance with the mock database
    a := &bapp.App{
        DB: testDB,
        R:  mux.NewRouter(),
    }

     //create a mock user to use for authentication
     user := models.User{
        ID: "81",
        Name:     "testuser",
        Email:    "testuser@example.com",
        Password: "testpassword",
    }
    
    body, _ := json.Marshal(user)
    req, err := http.NewRequest("POST", "/api/addUser", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    //mock request recorder
    mockRec := httptest.NewRecorder()

    //make request with mockRec and response
    a.R.HandleFunc("/api/addUser", a.AddUserHandler)
    a.R.ServeHTTP(mockRec, req)

    //check status code
    if status := mockRec.Code; status != http.StatusOK {
        t.Errorf("Wrong status returned, got %v want %v", status, http.StatusOK)
    }

    //check response body
    var responseUser models.User
    err = json.NewDecoder(mockRec.Body).Decode(&responseUser)
    if err != nil {
        t.Fatal(err)
    }
    if responseUser.ID != user.ID || responseUser.Name != user.Name || responseUser.Email != user.Email {
        t.Errorf("Returned unexpected user data, got %v, want %v", responseUser, user)
    }
}