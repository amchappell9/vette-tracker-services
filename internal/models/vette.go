package models

type Vette struct {
	ID               string   `json:"id"`
	Date             string   `json:"date"`
	UserID           string   `json:"userId"`
	Year             string   `json:"year"`
	Miles            string   `json:"miles"`
	Cost             string   `json:"cost"`
	TransmissionType string   `json:"transmissionType"`
	ExteriorColor    string   `json:"exteriorColor"`
	InteriorColor    string   `json:"interiorColor"`
	Submodel         string   `json:"submodel"`
	Trim             string   `json:"trim"`
	Packages         []string `json:"packages"`
	Link             string   `json:"link"`
}
