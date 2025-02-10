package main

import (
	"myapp/internal/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	//User endpoint

	// Add code here

	r.Run()
}
