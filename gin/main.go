package main

import (
	"github.com/HarshThakur1509/boilerplate/gin/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	//User endpoint

	// Add code here

	r.Run()
}
