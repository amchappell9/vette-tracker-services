package models

import (
	"time"

	"github.com/lib/pq"
)

type Vette struct {
	ID               int            `json:"id"`
	Date             time.Time      `json:"date"`
	UserID           string         `json:"userId"`
	Year             int16          `json:"year"`
	Miles            int            `json:"miles"`
	Cost             float64        `json:"cost"`
	TransmissionType string         `json:"transmissionType"`
	ExteriorColor    string         `json:"exteriorColor"`
	InteriorColor    string         `json:"interiorColor"`
	Submodel         string         `json:"submodel"`
	Trim             string         `json:"trim"`
	Packages         pq.StringArray `json:"packages"`
	Link             string         `json:"link"`
}
