package main

import (
	"vette-tracker-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.PingHandler)
	r.GET("/vette/:id", handlers.GetVetteHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
