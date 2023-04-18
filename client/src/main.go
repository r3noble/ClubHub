package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/bapp"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
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
	fmt.Println("ClubHub running, awaiting requests...")
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
	cdb, cerr := gorm.Open(sqlite.Open("clubs.db"), &gorm.Config{})
	if cerr != nil {
		panic("Error in opening CDB")
	}
	cerr = cdb.AutoMigrate(&models.Club{})
	if cerr != nil {
		panic("Error in migrating CDB")
	}
	edb, eerr := gorm.Open(sqlite.Open("clubs.db"), &gorm.Config{})
	if cerr != nil {
		panic("Error in opening CDB")
	}
	cerr = edb.AutoMigrate(&models.Event{})
	if cerr != nil {
		panic("Error in migrating EDB")
	}
	app := bapp.App{
		DB:  db,
		Cdb: cdb,
		R:   mux.NewRouter(),
		Edb: edb,
	}

	//hardcodes test user to db
	hardCoder := models.User{
		ID:       "123",
		Name:     "tester",
		Email:    "tester@example.com",
		Password: "password123",
		Clubs:    "WECE",
	}
	err = app.DB.Create(hardCoder).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
	}
	//app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}
	wece := models.Club{
		Name:      "WECE",
		President: "Jenna Sheldon",
		VP:        "Sarah Schultz",
		Treasurer: "Isabella Carmen",
		About:     "This is about promoting the inclusion of Women in the fields of both computer and electrical engineering",
	}
	ieee := models.Club{
		Name:      "IEEE",
		President: "Idk lol",
		VP:        "Idk lol",
		Treasurer: "Idk lol",
		About:     "Idk lol",
	}
	ufsit := models.Club{
		Name:      "UFSIT",
		President: "Idk lol",
		VP:        "Idk lol",
		Treasurer: "Idk lol",
		About:     "Idk lol",
	}
	wicse := models.Club{
		Name:      "WICSE",
		President: "Idk lol",
		VP:        "Idk lol",
		Treasurer: "Idk lol",
		About:     "Idk lol",
	}
	err = app.Cdb.Create(wece).Error
	if err != nil {
		fmt.Println("wece not added!")
	}
	err = app.Cdb.Create(ieee).Error
	if err != nil {
		fmt.Println("ieee not added!")
	}
	err = app.Cdb.Create(ufsit).Error
	if err != nil {
		fmt.Println("ufsit not added!")
	}
	err = app.Cdb.Create(wicse).Error
	if err != nil {
		fmt.Println("wicse not added!")
	}
	app.Start()
}
