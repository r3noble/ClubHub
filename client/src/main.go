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
	edb, cerr := gorm.Open(sqlite.Open("clubs.db"), &gorm.Config{})
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
	JShel := models.User{
		ID:       "228",
		Name:     "Jenna Sheldon",
		Email:    "jennasheldon@ufl.edu",
		Password: "password123",
		Clubs:    "WECE",
	}
	ICar := models.User{
		ID:       "333",
		Name:     "Isabella Cratem",
		Email:    "isabellacratem@ufl.edu",
		Password: "password123",
		Clubs:    "WECE",
	}
	SSchul := models.User{
		ID:       "111",
		Name:     "Sarah Schultz",
		Email:    "sarahschultz@ufl.edu",
		Password: "password123",
		Clubs:    "WECE",
	}
	err = app.DB.Create(hardCoder).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
	}
	err = app.DB.Create(JShel).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
	}
	err = app.DB.Create(ICar).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
	}
	err = app.DB.Create(SSchul).Error
	if err != nil {
		fmt.Println("Tester unsuccessfully hard-coded to db")
	}
	//app.u["Cole"] = User{ID: "1", Name: "Cole", Email: "cole@rottenberg.org", Password: "pass"}
	wece := models.Club{
		Name:      "WECE",
		President: "Jenna Sheldon",
		VP:        "Sarah Schultz",
		Treasurer: "Isabella Cratem",
		About:     "This is about promoting the inclusion of Women in the fields of both computer and electrical engineering",
	}
	ieee := models.Club{
		Name:      "IEEE",
		President: "Conrad Hellwege",
		VP:        "Julian Moldonado",
		Treasurer: "Justin Nagovskiy",
		About:     "The Institute of Electric and Electronics Engineers is the leading professional association for the advancement of technology. It is the world's largest technical society, bringing members access to the industry's most essential technical information, networking opportunities, career development tools, and many other exclusive benefits.",
	}
	ufsit := models.Club{
		Name:      "UFSIT",
		President: "Gabriella N",
		VP:        "Jon P",
		Treasurer: "Rachel O",
		About:     "UFSIT works to provide a welcoming environment for students to learn more about all areas of information security, including topics such as penetration testing, reverse engineering, vulnerability research, digital forensics, and more.",
	}
	wicse := models.Club{
		Name:      "WICSE",
		President: "Robin Fintz",
		VP:        "Minuet Greenberg",
		Treasurer: "Katja Karoleski",
		About:     "WiCSE is the official student ACM-W chapter at the University of Florida. They are dedicated to increasing the representation of women in Computer Science.",
	}
	sae := models.Club{
		Name:      "SAE",
		President: "Mitchell Thoeni",
		VP:        "Cam Lott",
		Treasurer: "Kyle Myott",
		About:     "Also known as Gator Motorsports, SAE offers a hands on experience designing, manufacturing, and marketing a technically advanced product: a small scale formula racecar.",
	}
	aiaa := models.Club{
		Name:      "AIAA",
		President: "Jose Aguilar",
		VP:        "Esha Shah",
		Treasurer: "Max Chern",
		About:     "AIAA promotes professional development for AE majors as well as ME majors with an interest in space and aeronautics through information sessions, tours, guest speakers, community outreach events, professional workshops, industry technical talks and other programming.",
	}
	acm := models.Club{
		Name:      "ACM",
		President: "Jennifer Lopez",
		VP:        "Matthew Hanson",
		Treasurer: "Joonho Jun",
		About:     "The University of Florida's Association for Computing Machinery works to promote computer science education through professional events, social events, and Special Interest Groups.",
	}
	dreamteam := models.Club{
		Name:      "dreamteam",
		President: "Emily Jones",
		VP:        "Neeva Sethi",
		Treasurer: "Vijayasai Somasundaram",
		About:     "Dream Team Engineering is a group of engineering and non-engineering students alike that are dedicated to making novel technologies to enhance the lives of the patients at UF Health Shands, the surrounding Gainesville community, and beyond.",
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
	err = app.Cdb.Create(sae).Error
	if err != nil {
		fmt.Println("sae not added!")
	}
	err = app.Cdb.Create(aiaa).Error
	if err != nil {
		fmt.Println("dreamteam not added!")
	}
	err = app.Cdb.Create(acm).Error
	if err != nil {
		fmt.Println("acm not added!")
	}
	err = app.Cdb.Create(dreamteam).Error
	if err != nil {
		fmt.Println("dreamteam not added!")
	}
	app.Start()
}
