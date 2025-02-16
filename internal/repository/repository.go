package repository

import (
	"database/sql"
	"vette-tracker-services/internal/models"
)

// This is the data access layer

type VetteRepositoryInterface interface {
	GetVettes() ([]models.Vette, error)
	GetVetteByID(id int) (models.Vette, error)
	InsertVette(vette models.Vette) (models.Vette, error)
	UpdateVette(vetteID int, vette models.Vette) (models.Vette, error)
	DeleteVette(vetteID int) error
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
		SELECT id, created_date, updated_date, deleted_date, user_id, year, 
					miles, cost, transmission_type, exterior_color, interior_color, 
			submodel, trim, packages, link
		FROM vettes
		WHERE deleted_date IS NULL
		ORDER BY updated_date desc 
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
			&v.CreatedDate,
			&v.UpdatedDate,
			&v.DeletedDate,
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
		SELECT id, created_date, updated_date, deleted_date, user_id, year, miles, cost, 
			transmission_type, exterior_color, interior_color, submodel, trim, packages, link
		FROM vettes
		WHERE id = $1 AND deleted_date IS NULL
	`, vetteID).Scan(
		&v.ID,
		&v.CreatedDate,
		&v.UpdatedDate,
		&v.DeletedDate,
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

func (r *VetteRepository) InsertVette(vette models.Vette) (models.Vette, error) {
	var insertedVette models.Vette

	err := r.db.QueryRow(`
        INSERT INTO vettes (
            created_date, updated_date, deleted_date, user_id, year, miles, cost, 
            transmission_type, exterior_color, interior_color, submodel, trim, 
            packages, link
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        RETURNING id, created_date, updated_date, deleted_date, user_id, year, 
            miles, cost, transmission_type, exterior_color, interior_color, 
            submodel, trim, packages, link
    `,
		vette.CreatedDate, vette.UpdatedDate, vette.DeletedDate, vette.UserID,
		vette.Year, vette.Miles, vette.Cost, vette.TransmissionType,
		vette.ExteriorColor, vette.InteriorColor, vette.Submodel, vette.Trim,
		vette.Packages, vette.Link).Scan(
		&insertedVette.ID, &insertedVette.CreatedDate, &insertedVette.UpdatedDate,
		&insertedVette.DeletedDate, &insertedVette.UserID, &insertedVette.Year,
		&insertedVette.Miles, &insertedVette.Cost, &insertedVette.TransmissionType,
		&insertedVette.ExteriorColor, &insertedVette.InteriorColor,
		&insertedVette.Submodel, &insertedVette.Trim, &insertedVette.Packages,
		&insertedVette.Link,
	)

	if err != nil {
		return models.Vette{}, err
	}

	return insertedVette, nil
}

func (r *VetteRepository) UpdateVette(vetteID int, vette models.Vette) (models.Vette, error) {
	var updatedVette models.Vette

	err := r.db.QueryRow(`
        UPDATE vettes SET 
            updated_date = $1,
            year = $2,
            miles = $3,
            cost = $4,
            transmission_type = $5,
            exterior_color = $6,
            interior_color = $7,
            submodel = $8,
            trim = $9,
            packages = $10,
            link = $11
        WHERE id = $12
        RETURNING id, created_date, updated_date, deleted_date, user_id, year, 
            miles, cost, transmission_type, exterior_color, interior_color, 
            submodel, trim, packages, link
    `,
		vette.UpdatedDate,
		vette.Year,
		vette.Miles,
		vette.Cost,
		vette.TransmissionType,
		vette.ExteriorColor,
		vette.InteriorColor,
		vette.Submodel,
		vette.Trim,
		vette.Packages,
		vette.Link,
		vetteID,
	).Scan(
		&updatedVette.ID,
		&updatedVette.CreatedDate,
		&updatedVette.UpdatedDate,
		&updatedVette.DeletedDate,
		&updatedVette.UserID,
		&updatedVette.Year,
		&updatedVette.Miles,
		&updatedVette.Cost,
		&updatedVette.TransmissionType,
		&updatedVette.ExteriorColor,
		&updatedVette.InteriorColor,
		&updatedVette.Submodel,
		&updatedVette.Trim,
		&updatedVette.Packages,
		&updatedVette.Link,
	)

	if err == sql.ErrNoRows {
		return models.Vette{}, sql.ErrNoRows
	}

	if err != nil {
		return models.Vette{}, err
	}

	return updatedVette, nil
}

func (r *VetteRepository) GetVettesCount() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM vettes").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
