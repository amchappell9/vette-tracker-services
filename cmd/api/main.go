package main

import (
	"log"
	"vette-tracker-services/internal/database"
	"vette-tracker-services/internal/handlers"
	"vette-tracker-services/internal/repository"
	"vette-tracker-services/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewConnection()

	if err != nil {
		log.Fatalf("Could not initialize database connection: %v", err)
	}

	defer db.Close()

	// Initialize layers
	vetteRepo := repository.NewVetteRepository(db)
	vetteService := service.NewVetteService(vetteRepo)
	handler := handlers.NewHandler(vetteService)

	r := gin.Default()
	r.GET("/ping", handler.PingHandler)
	r.GET("/vettes", handler.GetVettesHandler)
	r.GET("/vettes/:id", handler.GetVetteHandler)
	r.POST("/vettes", handler.CreateVetteHandler)
	r.GET("/vettes/count", handler.GetVetteCountHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
