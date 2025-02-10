package main

import (
	"myapp/internal/initializers"
	"myapp/internal/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

	// Add code here

}

func main() {
	server := routes.NewApiServer(":3000")
	server.Run()
}
