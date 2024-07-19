package models

import (
	"database/sql"
	"fmt"
)

type EmailToInvite struct {
	ID     string `json:"id"`
	TripId string `json:"trip_id"`
	Email  string `json:"email"`
}

type EmailsRepository struct {
	DB *sql.DB
}

// Inicia o repositório com DB
func NewEmailsRepository(db *sql.DB) *EmailsRepository {
	return &EmailsRepository{DB: db}
}

// Convida pessoas para uma viagem
func (repo *EmailsRepository) EmailsToInvite(email *EmailToInvite) error {
	// Query SQL
	query := `
		INSERT INTO emails_to_invite
		    (id, trip_id, email)
		VALUES
			($1, $2, $3)
	`

	fmt.Println(email)

	_, err := repo.DB.Exec(
		query,
		email.ID,
		email.TripId,
		email.Email,
	)

	// Erro ao criar Viagem no DB
	if err != nil {
		return fmt.Errorf("erro ao convidar emails para viagem: %w", err)
	}

	return nil
}
