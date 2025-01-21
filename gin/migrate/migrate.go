package main

import (
	"log"

	"github.com/HarshThakur1509/boilerplate/gin/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	log.Println("Starting database migrations...")

	// Add code here

	err := initializers.DB.AutoMigrate()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully!")

}
