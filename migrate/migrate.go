package main

import (
	"boilerplate/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {

	// Add code here

	initializers.DB.AutoMigrate()
}
