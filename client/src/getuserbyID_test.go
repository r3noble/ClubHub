package main

import (
    "testing"
)

func TestGetUserByID(t *testing.T) {
    // create a new app
    app := &App{
        u: map[string]User{
            "1": {ID: 1, Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"},
        },
    }

    // call GetUserByID with an existing user ID
    got, err := app.GetUserByID("1")

    // create a User struct that resembles the expected result
    want := &User{ID: 1, Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

    // check if the returned User and error match the expected values
    if got == nil || *got != *want || err != nil {
        t.Errorf("got %v, want %v, error %v", got, want, err)
    }

    // call GetUserByID with a non-existing user ID
    got, err = app.GetUserByID("2")

    // check if nil and an error message are returned
    if got != nil || err == nil || err.Error() != "user with ID 2 not found" {
        t.Errorf("got %v, want nil, error %v", got, err)
    }
}

