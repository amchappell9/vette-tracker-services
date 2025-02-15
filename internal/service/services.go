package service

import (
	"time"
	"vette-tracker-services/internal/models"
	"vette-tracker-services/internal/repository"
)

// The is the business logic layer

type VetteServiceInterface interface {
	GetVettes() ([]models.Vette, error)
	GetVette(id int) (models.Vette, error)
	GetVettesCount() (int, error)
	CreateVette(vette models.Vette) (models.Vette, error)
}

type VetteService struct {
	repo *repository.VetteRepository
}

func NewVetteService(repo *repository.VetteRepository) *VetteService {
	return &VetteService{repo: repo}
}

func (s *VetteService) GetVettes() ([]models.Vette, error) {
	// TODO: business logic such as filtering, authorization
	return s.repo.GetVettes()
}

func (s *VetteService) GetVette(vetteID int) (models.Vette, error) {
	return s.repo.GetVetteByID(vetteID)
}

func (s *VetteService) GetVettesCount() (int, error) {
	return s.repo.GetVettesCount()
}

func (s *VetteService) CreateVette(createVetteReq models.CreateVetteRequest) (models.Vette, error) {
	vette := models.Vette{
		CreatedDate:      time.Now(),
		UpdatedDate:      time.Now(),
		UserID:           "123", // Placeholder user ID
		Year:             createVetteReq.Year,
		Miles:            createVetteReq.Miles,
		Cost:             createVetteReq.Cost,
		TransmissionType: createVetteReq.TransmissionType,
		ExteriorColor:    createVetteReq.ExteriorColor,
		InteriorColor:    createVetteReq.InteriorColor,
		Submodel:         createVetteReq.Submodel,
		Trim:             createVetteReq.Trim,
		Packages:         createVetteReq.Packages,
		Link:             createVetteReq.Link,
	}

	return s.repo.InsertVette(vette)
}
