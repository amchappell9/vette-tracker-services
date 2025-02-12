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

	handler := handlers.NewHandler(db)

	r := gin.Default()
	r.GET("/ping", handler.PingHandler)
	r.GET("/vette", handler.GetVettes)
	r.GET("/vette/:id", handler.GetVetteHandler)
	r.GET("/count", handler.GetVetteCountHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
