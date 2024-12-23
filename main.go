package main

import (
	"github.com/HarshThakur1509/boilerplate/api"
	"github.com/HarshThakur1509/boilerplate/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	server := api.NewApiServer(":3000")
	server.Run()
}
