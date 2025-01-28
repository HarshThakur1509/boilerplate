package main

import (
	"log"

	"github.com/HarshThakur1509/boilerplate/standard/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
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
