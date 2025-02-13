package repository

import (
	"database/sql"
	"vette-tracker-services/internal/models"
)

// This is the data access layer

type VetteRepositoryInterface interface {
	GetVettes() ([]models.Vette, error)
	GetVetteByID(id int) (models.Vette, error)
	GetVettesCount() (int, error)
}

type VetteRepository struct {
	db *sql.DB
}

func NewVetteRepository(db *sql.DB) *VetteRepository {
	return &VetteRepository{db: db}
}

func (r *VetteRepository) GetVettes() ([]models.Vette, error) {
	rows, err := r.db.Query(`
        SELECT id, date, user_id, year, miles, cost, 
            transmission_type, exterior_color, interior_color, 
            submodel, trim, packages, link
        FROM vettes
        ORDER BY date desc 
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vettes []models.Vette
	for rows.Next() {
		var v models.Vette
		err := rows.Scan(
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
		if err != nil {
			return nil, err
		}
		vettes = append(vettes, v)
	}
	return vettes, rows.Err()
}

func (r *VetteRepository) GetVetteByID(vetteID int) (models.Vette, error) {
	var v models.Vette

	err := r.db.QueryRow(`
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

	// TODO: Improve this
	if err == sql.ErrNoRows {
		return models.Vette{}, err
	}

	if err != nil {
		return models.Vette{}, err
	}

	return v, nil
}

func (r *VetteRepository) GetVettesCount() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM vettes").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
