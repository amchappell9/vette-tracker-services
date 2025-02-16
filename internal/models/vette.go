package models

import (
	"time"

	"github.com/lib/pq"
)

type Vette struct {
	ID               int            `json:"id"`
	CreatedDate      time.Time      `json:"createdDate"`
	UpdatedDate      time.Time      `json:"updatedDate"`
	DeletedDate      *time.Time     `json:"-"`
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

/*
The structure a caller should pass when inserting a new request.
*/
type VetteRequestObj struct {
	Year             int16          `json:"year" binding:"required"`
	Miles            int            `json:"miles" binding:"required"`
	Cost             float64        `json:"cost" binding:"required"`
	TransmissionType string         `json:"transmissionType" binding:"required"`
	ExteriorColor    string         `json:"exteriorColor" binding:"required"`
	InteriorColor    string         `json:"interiorColor" binding:"required"`
	Submodel         string         `json:"submodel" binding:"required"`
	Trim             string         `json:"trim" binding:"required"`
	Packages         pq.StringArray `json:"packages"`
	Link             string         `json:"link"`
}
