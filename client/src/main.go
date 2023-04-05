package main

import (
	//"encoding/json"
	"fmt"
	//"math/rand"
	"net/http"
	//"strconv"
	//"sync"

	"github.com/gorilla/mux"
	//"github.com/rs/cors"
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
	app := App{
		db: db,
		u:  make(map[string]User),
		r:  mux.NewRouter(),
	}

	//hardcodes test user to db
	hardCoder := User{
		ID:       "123",
		Name:     "tester",
		Email:    "tester@example.com",
		Password: "password123",
	}
	err = app.db.Create(hardCoder).Error
	if err != nil {
		fmt.Println("Hardcoder unsuccessfully added to db")
		return
	}
	//app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}

	app.start()
}

