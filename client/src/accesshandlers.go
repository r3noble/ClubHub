package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"context"

	"github.com/jackc/pgx/v4"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &models.User{}
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
	//establish connection to the database (DB)
	DB, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/mydb")
	if err != nil {
		fmt.Println("Could not establish database connection")
		return
	}
	defer DB.Close(context.Background())
	//insert username, hashed pword, and userlevel into db
	insertData := `INSERT INTO USER (Name, Email, Password, userlevel)
					VALUES ($1, $2, $3, $4)`

	if _, err = DB.Exec(context.Background(), insertData, creds.Name, creds.Email, hashedPassword, creds.userLevel); err != nil {
		//if error return service error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Properly signed up if no errors til this point
}

/*func Signin(w http.ResponseWriter, r *http.Request) {
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
}*/

func SignIn(w http.ResponseWriter, r *http.Request) {
	creds := &models.User{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If the request body is not valid JSON, return a 400 Bad Request error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Establish a connection to the database
	db, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/mydb")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close(context.Background())

	// Retrieve the user with the specified email from the database
	row := db.QueryRow(context.Background(), "SELECT name, email, password, userLevel FROM users WHERE email = $1", creds.Email)
	var user models.User
	err = row.Scan(&user.Name, &user.Email, &user.Password, &user.userLevel)
	if err != nil {
		// If there is an error retrieving the user from the database, return a 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Use bcrypt to compare the entered password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		// If the passwords do not match, return a 401 Unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Return a success message in the response body
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sign in successful"))
}
