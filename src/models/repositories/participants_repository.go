package models

import (
	"database/sql"
	"fmt"
)

type Participants struct {
	ID               string `json:"id"`
	TripId           string `json:"-"`
	EmailsToInviteID string `json:"-"`
	Name             string `json:"name"`
	Email            string `json:"email"`
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

func (repo *ParticipantsRepository) FindParticipantsFromTrip(tripID string) ([]Participants, error) {
	query := `
		SELECT p.id, p.trip_id, p.emails_to_invite_id, p.name, e.email, p.is_confirmed
		FROM participants as p
		JOIN emails_to_invite as e ON e.id = p.emails_to_invite_id
		WHERE p.trip_id = $1
	`

	rows, err := repo.DB.Query(query, tripID)
	if err != nil {
		return nil, fmt.Errorf("erro ao encontrar links: %w", err)
	}
	defer rows.Close()

	var participants []Participants
	for rows.Next() {
		var participant Participants

		if err := rows.Scan(&participant.ID, &participant.TripId, &participant.EmailsToInviteID, &participant.Name, &participant.Email, &participant.IsConfirmed); err != nil {
			return nil, fmt.Errorf("erro ao escanear participants: %w", err)
		}

		participants = append(participants, participant)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %w", err)
	}

	return participants, nil
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
