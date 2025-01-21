package main

import "github.com/HarshThakur1509/boilerplate/standard/initializers"

func init() {
	initializers.ConnectDB()
}

func main() {

	// Add code here

	initializers.DB.AutoMigrate()
}
