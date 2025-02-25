package handlers

import (
	"log"
	"net/http"
	"strconv"
	"vette-tracker-services/internal/errors"
	"vette-tracker-services/internal/models"
	"vette-tracker-services/internal/service"

	"github.com/gin-gonic/gin"
)

type VetteHandlerInterface interface {
	GetVettes(c *gin.Context)
	GetVetteHandler(c *gin.Context)
	GetVetteCountHandler(c *gin.Context)
	CreateVetteHandler(c *gin.Context)
	UpdateVetteHandler(c *gin.Context)
	DeleteVette(c *gin.Context)
}

type Handler struct {
	vetteService service.VetteServiceInterface
}

func NewHandler(service service.VetteServiceInterface) *Handler {
	return &Handler{vetteService: service}
}

func (h *Handler) GetVettesHandler(c *gin.Context) {
	vettes, err := h.vetteService.GetVettes()

	if err != nil {
		log.Printf("Error getting vettes: %v\n", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, vettes)
}

func (h *Handler) GetVetteHandler(c *gin.Context) {
	id := c.Param("id")

	// Validate ID is passed and is numeric
	if id == "" {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID is required",
		})
		return
	}

	vetteID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID must be a numeric value",
		})
		return
	}

	vette, err := h.vetteService.GetVette(int(vetteID))

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, vette)
}

func (h *Handler) CreateVetteHandler(c *gin.Context) {
	var createRequestVette models.VetteRequestObj
	if err := c.ShouldBindJSON(&createRequestVette); err != nil {
		c.Error(&errors.ValidationError{
			Field:   "body",
			Message: "Invalid request body",
		})
		return
	}

	createdVette, err := h.vetteService.CreateVette(createRequestVette)

	if err != nil {
		c.Error(&errors.DatabaseError{
			Operation: "create_vette",
			Err:       err,
		})
		return
	}

	c.JSON(http.StatusCreated, createdVette)
}

func (h *Handler) UpdateVetteHandler(c *gin.Context) {
	id := c.Param("id")

	// Validate ID is passed and is numeric
	if id == "" {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID is required",
		})
		return
	}

	vetteID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID must be a numeric value",
		})
		return
	}

	var updateRequestVette models.VetteRequestObj
	if err := c.ShouldBindJSON(&updateRequestVette); err != nil {
		c.Error(&errors.ValidationError{
			Field:   "body",
			Message: "Invalid request body",
		})
		return
	}

	updatedVette, err := h.vetteService.UpdateVette(int(vetteID), updateRequestVette)
	if err != nil {
		c.Error(&errors.DatabaseError{
			Operation: "update_vette",
			Err:       err,
		})
		return
	}

	c.JSON(http.StatusOK, updatedVette)
}

func (h *Handler) DeleteVette(c *gin.Context) {
	id := c.Param("id")

	// Check that there's an ID
	if id == "" {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID is required",
		})
		return
	}

	// Check that ID is valid
	vetteID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Error(&errors.ValidationError{
			Field:   "id",
			Message: "ID must be a numeric value",
		})
		return
	}

	err = h.vetteService.DeleteVette(int(vetteID))
	if err != nil {
		c.Error(&errors.DatabaseError{
			Operation: "delete_vette",
			Err:       err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *Handler) GetVetteCountHandler(c *gin.Context) {
	count, err := h.vetteService.GetVettesCount()
	if err != nil {
		c.Error(&errors.DatabaseError{
			Operation: "get_vette_count",
			Err:       err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
