package main

import (
	"boilerplate/api"
	"boilerplate/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	server := api.NewApiServer(":3000")
	server.Run()
}
