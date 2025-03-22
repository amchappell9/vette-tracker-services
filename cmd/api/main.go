package main

import (
	"log"
	"os"
	"vette-tracker-services/internal/database"
	"vette-tracker-services/internal/handlers"
	"vette-tracker-services/internal/middleware"
	"vette-tracker-services/internal/repository"
	"vette-tracker-services/internal/service"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	key := os.Getenv("CLERK_SECRET_KEY")
	if key == "" {
		log.Fatal("CLERK_SECRET_KEY environment variable is not set")
	}

	// Create a new Clerk client instance
	clerk.SetKey(key)

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
	r.Use(middleware.ErrorHandler())

	// Utils (Public)
	r.GET("/ping", handlers.PingHandler)

	// Vette Handlers (Protected)
	protected := r.Group("")
	protected.Use(middleware.ClerkAuth())
	{
		protected.GET("/vettes", handler.GetVettesHandler)
		protected.GET("/vettes/:id", handler.GetVetteHandler)
		protected.POST("/vettes", handler.CreateVetteHandler)
		protected.PUT("/vettes/:id", handler.UpdateVetteHandler)
		protected.DELETE("/vettes/:id", handler.DeleteVette)
		protected.GET("/vettes/count", handler.GetVetteCountHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
