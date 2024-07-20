package models

import (
	"database/sql"
	"fmt"
)

type Participants struct {
	ID               string `json:"id"`
	TripId           string `json:"trip_id"`
	EmailsToInviteID string `json:"emails_to_invite_id"`
	Name             string `json:"name"`
	IsConfirmed      int    `json:"is_confirmed"`
}

type ParticipantsRepository struct {
	DB *sql.DB
}

func NewParticipantsRepository(db *sql.DB) *ParticipantsRepository {
	return &ParticipantsRepository{DB: db}
}

func (repo *ParticipantsRepository) CreateParticipants(participant *Participants) error {
	query := `
        INSERT INTO participants
            (id, trip_id, emails_to_invite_id, name, is_confirmed)
        VALUES
            ($1, $2, $3, $4, $5)
    `

	_, err := repo.DB.Exec(
		query,
		participant.ID,
		participant.TripId,
		participant.EmailsToInviteID,
		participant.Name,
		participant.IsConfirmed,
	)

	if err != nil {
		return fmt.Errorf("erro ao criar participantes: %w", err)
	}

	return nil
}

func (repo *ParticipantsRepository) FindParticipantsFromTrip(tripID string) (*Participants, error) {
	query := `
		SELECT p.id, p.name, p.is_confirmed, e.email
        FROM participants as p
        JOIN emails_to_invite as e
		ON e.id = p.emails_to_invite_id
        WHERE p.trip_id = $1
	`
	var participant Participants
	row := repo.DB.QueryRow(query, tripID)

	if err := row.Scan(&participant.ID, &participant.TripId, &participant.EmailsToInviteID, &participant.Name, &participant.IsConfirmed); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("viagem ou participante não encontrada com ID da viagem: %s", tripID)
		}
		return nil, fmt.Errorf("erro ao encontrar participante: %w", err)
	}

	return &participant, nil
}

func (repo *ParticipantsRepository) UpdateParticipant(participantID string) error {
	// Query SQL
	query := `
		UPDATE 
			participants
		SET
			is_confirmed = 1
		WHERE
			id = $1
	`

	_, err := repo.DB.Exec(
		query,
		participantID,
	)

	// Erro ao atualizar Viagem no DB
	if err != nil {
		return fmt.Errorf("erro ao atualizar participante: %w", err)
	}

	return nil
}
