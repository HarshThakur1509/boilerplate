package main

import (
	"log"

	"github.com/HarshThakur1509/boilerplate/standard/initializers"
	"github.com/HarshThakur1509/boilerplate/standard/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {

	log.Println("Starting database migrations...")

	User := &models.User{}
	// Add code here

	err := initializers.DB.AutoMigrate(User)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully!")
}
