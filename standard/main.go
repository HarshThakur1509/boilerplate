package main

import (
	"github.com/HarshThakur1509/boilerplate/standard/api"
	"github.com/HarshThakur1509/boilerplate/standard/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
	server := api.NewApiServer(":3000")
	server.Run()
}
