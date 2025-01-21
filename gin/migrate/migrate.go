package main

import "github.com/HarshThakur1509/boilerplate/gin/initializers"

func init() {
	initializers.ConnectToDB()
}

func main() {
	// Add code here
	initializers.DB.AutoMigrate()

}
