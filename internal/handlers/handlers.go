package handlers

import (
	"database/sql"
	"log"
	"strconv"
	"vette-tracker-services/internal/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func (h *Handler) GetVettes(c *gin.Context) {
	// TODO: validate user is authenticated

	// TODO: Add where clause for user id
	rows, err := h.db.Query(`
		SELECT id, date, user_id, year, miles, cost, 
			transmission_type, exterior_color, interior_color, 
			submodel, trim, packages, link
		FROM vettes
		ORDER BY date desc 
	`)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve vettes"})
	}

	defer rows.Close()

	var vettes []models.Vette

	for rows.Next() {
		var v models.Vette

		err := rows.Scan(&v.ID,
			&v.Date,
			&v.UserID,
			&v.Year,
			&v.Miles,
			&v.Cost,
			&v.TransmissionType,
			&v.ExteriorColor,
			&v.InteriorColor,
			&v.Submodel,
			&v.Trim,
			&v.Packages,
			&v.Link,
		)

		if err != nil {
			log.Printf("Error scanning vette: %v\n", err)
			c.JSON(500, gin.H{"error": "Failed to parse vettes"})
			return
		}

		vettes = append(vettes, v)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v\n", err)
		c.JSON(500, gin.H{"error": "Error iterating over rows"})
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

	var v models.Vette

	err = h.db.QueryRow(`
		SELECT id, date, user_id, year, miles, cost, transmission_type, exterior_color,
			interior_color, submodel, trim, packages, link
		FROM vettes
		WHERE id = $1
	`, vetteID).Scan(
		&v.ID,
		&v.Date,
		&v.UserID,
		&v.Year,
		&v.Miles,
		&v.Cost,
		&v.TransmissionType,
		&v.ExteriorColor,
		&v.InteriorColor,
		&v.Submodel,
		&v.Trim,
		&v.Packages,
		&v.Link,
	)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "Vette not found"})
		return
	}

	if err != nil {
		log.Printf("Failed to retrieve vette: %v\n", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve vette"})
		return
	}

	c.JSON(200, v)
}

func (h *Handler) GetVetteCountHandler(c *gin.Context) {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM vettes").Scan(&count)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get count"})
		return
	}

	c.JSON(200, gin.H{"count": count})
}
