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
	CreateVette(vette models.VetteRequestObj) (models.Vette, error)
	UpdateVette(vetteId int, updateVetteReq models.VetteRequestObj) (models.Vette, error)
	DeleteVette(vetteId int) error
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

func (s *VetteService) CreateVette(createVetteReq models.VetteRequestObj) (models.Vette, error) {
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

func (s *VetteService) UpdateVette(vetteID int, updateVetteReq models.VetteRequestObj) (models.Vette, error) {
	// First get the existing vette to preserve unchangeable fields
	existingVette, err := s.repo.GetVetteByID(vetteID)
	if err != nil {
		return models.Vette{}, err
	}

	// Create updated vette model preserving some original fields
	vette := models.Vette{
		ID:               existingVette.ID,
		CreatedDate:      existingVette.CreatedDate,
		UpdatedDate:      time.Now(),
		DeletedDate:      existingVette.DeletedDate,
		UserID:           existingVette.UserID,
		Year:             updateVetteReq.Year,
		Miles:            updateVetteReq.Miles,
		Cost:             updateVetteReq.Cost,
		TransmissionType: updateVetteReq.TransmissionType,
		ExteriorColor:    updateVetteReq.ExteriorColor,
		InteriorColor:    updateVetteReq.InteriorColor,
		Submodel:         updateVetteReq.Submodel,
		Trim:             updateVetteReq.Trim,
		Packages:         updateVetteReq.Packages,
		Link:             updateVetteReq.Link,
	}

	return s.repo.UpdateVette(vetteID, vette)
}

func (s *VetteService) DeleteVette(vetteID int) error {
	existingVette, err := s.repo.GetVetteByID(vetteID)

	if err != nil {
		return err
	}

	now := time.Now()

	deletedVette := models.Vette{
		ID:               existingVette.ID,
		CreatedDate:      existingVette.CreatedDate,
		UpdatedDate:      existingVette.UpdatedDate,
		DeletedDate:      &now,
		UserID:           existingVette.UserID,
		Year:             existingVette.Year,
		Miles:            existingVette.Miles,
		Cost:             existingVette.Cost,
		TransmissionType: existingVette.TransmissionType,
		ExteriorColor:    existingVette.ExteriorColor,
		InteriorColor:    existingVette.InteriorColor,
		Submodel:         existingVette.Submodel,
		Trim:             existingVette.Trim,
		Packages:         existingVette.Packages,
		Link:             existingVette.Link,
	}

	_, err = s.repo.UpdateVette(vetteID, deletedVette)

	if err != nil {
		return err
	}

	return nil
}
