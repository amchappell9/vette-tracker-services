package main

import (
	"log"
	"vette-tracker-services/internal/database"
	"vette-tracker-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewConnection()

	if err != nil {
		log.Fatalf("Could not initialize database connection: %v", err)
	}

	defer db.Close()

	r := gin.Default()
	r.GET("/ping", handlers.PingHandler)
	r.GET("/vette/:id", handlers.GetVetteHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
