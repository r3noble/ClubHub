package main

import (
	"testing"
)

func (a *App) TestGetUserByID(t *testing.T) {
	//what is actually returned
	got, err := a.GetUserByID("1")
	//create user struct to resemble what we wanted to be return
	want := &User{ID: 1, Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

	if got != want || err != nil{
        t.Errorf("got %q, wanted %q", got, want)
    }
}

/*func (a *App) TestloginHandler(t *testing.T) {
	//create mock request
	body := []byte(`{"username": "testuser", "password": "testpass"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	//create mock response recorder
	mockRec := httptest.NewRecorder()

	//call loginHandler with mock request and recorder
	handler := http.HandleFunc("/user/login", a.loginHandler)
	handler.ServeHTTP(mockRec, req)

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