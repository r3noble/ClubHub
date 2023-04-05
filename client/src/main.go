package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
)


func WriteOnceMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{w, false}
		h.ServeHTTP(wrappedWriter, r)
		if !wrappedWriter.wroteHeader {
			wrappedWriter.WriteHeader(http.StatusOK)
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.wroteHeader = true
}

func main() {
	//Initialize and open DB here
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Error in opening DB")
	}
	//calls AutoMigrate and throws error if cannot migrate
	//formats db to replicate user struct
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error in migrating db")
	}
	app := bapp.App{
		DB: db,
		R: mux.NewRouter(),
	}

	//hardcodes test user to db
	hardCoder := models.User{
		ID:       "123",
		Name:     "tester",
		Email:    "tester@example.com",
		Password: "password123",
	}
	err = app.DB.Create(hardCoder).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
		return
	}
	//app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

	app.Start()
}

