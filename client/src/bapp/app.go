package bapp

import(
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"	
)


// var userMap map[int]User
type App struct {
	//DB is for user storage
	DB *gorm.DB
	//Cdb is for club storage
	Cdb *gorm.DB
	//Edb is for event storage
	Edb *gorm.DB
	R  *mux.Router
	mu sync.Mutex
}

func (a *App) Start() {

	a.R.Use(func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	a.R.HandleFunc("/api/health", HealthCheck).Methods("GET")
	//query-based matching using id
	//user CRUD APIs
	a.R.HandleFunc("/api/getUser/{id}", a.IdHandler).Methods("GET")
	a.R.HandleFunc("/api/addUser", a.AddUserHandler).Methods("POST")
	a.R.HandleFunc("/api/login", a.LoginHandler).Methods("POST") // handlers login
	//club CRUD APIs
	a.R.HandleFunc("/api/addClub", a.AddClubHandler).Methods("POST")
	a.R.HandleFunc("/api/joinClub", a.JoinClubHandler).Methods("POST")
	//a.r.HandleFunc("/api/getClub", a.GetClubHandler).Methods("GET")
	//a.r.HandleFunc("/api/delClub", a.DeleteClubHandler).Methods("DELETE")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(a.R)

	http.ListenAndServe(":8080", handler)
}