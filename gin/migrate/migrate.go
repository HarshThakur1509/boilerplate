package main

import "github.com/HarshThakur1509/boilerplate/gin/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Add code here
	initializers.DB.AutoMigrate()

}
