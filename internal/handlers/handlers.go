package handlers

import (
	"strconv"
	"vette-tracker-services/internal/models"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func GetVetteHandler(c *gin.Context) {
	id := c.Param("id")

	// Validate ID is passed and is numeric
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID must be a numeric value"})
		return
	}

	vette := models.Vette{
		ID:               id,
		Date:             "11-06-2024",
		UserID:           "user_2QfpZzBgJPSVcvXnnEike2m5Dvk",
		Year:             "2019",
		Miles:            "295",
		Cost:             "68000",
		TransmissionType: "Automatic",
		ExteriorColor:    "Torch Red",
		InteriorColor:    "Red",
		Submodel:         "Grand Sport",
		Trim:             "2LT",
		Packages:         []string{"MRC", "NPP", "PDR"},
		Link:             "https://www.corvetteforum.com/forums/c7-corvettes-for-sale/4876333-2019-grand-sport-convertible-only-295-miles.html",
	}

	c.JSON(200, vette)
}
