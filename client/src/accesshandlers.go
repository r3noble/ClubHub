package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// models struct of users, replace later
type Credentials struct {
	Password  string `json:"password", db:"password"`
	Username  string `json:"username", db:"username"`
	UserLevel string `json:"userlevel, db:"userlevel"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		//if issue occurs throw 400 bad request error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//use bcrypt to salt & hash pword
	//8 is hashing value
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Could not hash password")
		return
	}

	//insert username, hashed pword, and userlevel into db
	insertData := `INSERT INTO users (name, password, userlevel)
					VALUES ($1, $2, $3)`

	if _, err = db.Query(insertData, creds.Username, hashedPassword, creds.UserLevel); err != nil {
		//if error return service error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Properly signed up if no errors til this point
}

func Signin(w http.ResponseWriter, r *http.Request) {
	//create new credentials instance from parsed request
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//get pword found in DB for username input
	result := db.QueryRow("SELECT userlevel, password FROM users WHERE username=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//create another credentials instance to store what we retrieve from the db
	stored := &Credentials{}
	//store obtained pword
	err = result.Scan(&stored.Password, &stored.UserLevel)
	if err != nil {
		if err == sql.ErrNoRows {
			//invalid password entered => **DO SOMETHING****
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//other issue, do something else
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(creds.Password)); err != nil {
		//***** If the two passwords don't match, return a 401 status and DO SOMETHING *****
		w.WriteHeader(http.StatusUnauthorized)
	}
}
