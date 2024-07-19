package models

import (
	"database/sql"
	"fmt"
)

type Trip struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	OwnerName   string `json:"owner_name"`
	OwnerEmail  string `json:"owner_email"`
	Status      int    `json:"status"`
}

type TripRepository struct {
	DB *sql.DB
}

// Inicia o repositório com DB
func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{DB: db}
}

// Cria uma nova viagem no DB
func (repo *TripRepository) CreateTrip(trip *Trip) error {
	// Query SQL
	query := `
		INSERT INTO trips
			(id, destination, start_date, end_date, owner_name, owner_email, status)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := repo.DB.Exec(
		query,
		trip.ID,
		trip.Destination,
		trip.StartDate,
		trip.EndDate,
		trip.OwnerName,
		trip.OwnerEmail,
		trip.Status,
	)

	// Erro ao criar Viagem no DB
	if err != nil {
		return fmt.Errorf("erro ao criar viagem: %w", err)
	}

	return nil
}

// FindTripByID recupera uma viagem com base no ID
func (repo *TripRepository) FindTripByID(tripID string) (*Trip, error) {
	query := `
		SELECT * FROM trips WHERE id = $1
	`

	var trip Trip
	row := repo.DB.QueryRow(query, tripID)

	if err := row.Scan(&trip.ID, &trip.Destination, &trip.StartDate, &trip.EndDate, &trip.OwnerName, &trip.OwnerEmail, &trip.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("viagem não encontrada com ID: %s", tripID)
		}
		return nil, fmt.Errorf("erro ao encontrar viagem: %w", err)
	}

	return &trip, nil
}

func (repo *TripRepository) UpdateTrip(trip_id string) error {
	// Query SQL
	query := `
		UPDATE 
			trips
		SET
			status = 1
		WHERE
			id = $1
	`

	_, err := repo.DB.Exec(
		query,
		trip_id,
	)

	// Erro ao atualizar Viagem no DB
	if err != nil {
		return fmt.Errorf("erro ao atualizar viagem: %w", err)
	}

	return nil
}
