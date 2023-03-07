package main

import(
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
)

func TestLoginHandler(t *testing.T) {
    a := &App{
        u: map[string]User{
            "1": {ID: 1, Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"},
        },
    }
    //create mock request
    body := []byte(`{"username": "testuser", "password": "testpass"}`)
    req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    //create mock response recorder
    mockRec := httptest.NewRecorder()

    //call loginHandler with mock request and recorder
    http.HandleFunc("/user/login", a.loginHandler)
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

