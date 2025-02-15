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
        SELECT id, created_date, updated_date, deleted_date, user_id, year, 
					miles, cost, transmission_type, exterior_color, interior_color, 
          submodel, trim, packages, link
        FROM vettes
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
		WHERE id = $1
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

func (r *VetteRepository) GetVettesCount() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM vettes").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
