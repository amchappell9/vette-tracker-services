package handlers

import (
	"log"
	"strconv"
	"vette-tracker-services/internal/models"
	"vette-tracker-services/internal/service"

	"github.com/gin-gonic/gin"
)

type VetteHandlerInterface interface {
	PingHandler(c *gin.Context)
	GetVettes(c *gin.Context)
	GetVetteHandler(c *gin.Context)
	GetVetteCountHandler(c *gin.Context)
}

type Handler struct {
	vetteService *service.VetteService
}

func NewHandler(service *service.VetteService) *Handler {
	return &Handler{vetteService: service}
}

func (h *Handler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func (h *Handler) GetVettesHandler(c *gin.Context) {
	vettes, err := h.vetteService.GetVettes()

	if err != nil {
		// TODO: Do I want to panic here?
		log.Printf("Error getting vettes: %v\n", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, vettes)
}

func (h *Handler) GetVetteHandler(c *gin.Context) {
	id := c.Param("id")

	// Validate ID is passed and is numeric
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	vetteID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID must be a numeric value"})
		return
	}

	vette, err := h.vetteService.GetVette(int(vetteID))

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
	}

	c.JSON(200, vette)
}

func (h *Handler) CreateVetteHandler(c *gin.Context) {
	var createRequestVette models.VetteRequestObj
	if err := c.ShouldBindJSON(&createRequestVette); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	createdVette, err := h.vetteService.CreateVette(createRequestVette)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create vette"})
		return
	}

	c.JSON(201, createdVette)
}

func (h *Handler) UpdateVetteHandler(c *gin.Context) {
	id := c.Param("id")

	// Validate ID is passed and is numeric
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	vetteID, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID must be a numeric value"})
		return
	}

	var updateRequestVette models.VetteRequestObj
	if err := c.ShouldBindJSON(&updateRequestVette); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedVette, err := h.vetteService.UpdateVette(int(vetteID), updateRequestVette)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update vette"})
		return
	}

	c.JSON(200, updatedVette)
}

func (h *Handler) GetVetteCountHandler(c *gin.Context) {
	count, err := h.vetteService.GetVettesCount()

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get count"})
		return
	}

	c.JSON(200, gin.H{"count": count})
}
