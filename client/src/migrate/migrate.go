package main

import (
	"fmt"
	"log"

	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/initializers"
	"github.com/r3noble/CEN3031-Project-Group/tree/main/client/src/models"
)

// creates connection pool to postgres db
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
