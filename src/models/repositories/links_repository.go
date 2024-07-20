package models

import (
	"database/sql"
	"fmt"
)

type Links struct {
	ID     string `json:"id"`
	TripID string `json:"trip_id"`
	Link   string `json:"link"`
	Title  string `json:"title"`
}

type LinksRepository struct {
	DB *sql.DB
}

// Inicia o repositório com DB
func NewLinksRepository(db *sql.DB) *LinksRepository {
	return &LinksRepository{DB: db}
}

func (repo *LinksRepository) CreateLink(link *Links) error {
	query := `
		INSERT INTO links
			(id, trip_id, link, title)
		VALUES
			($1, $2, $3, $4)
	`

	_, err := repo.DB.Exec(
		query,
		link.ID,
		link.TripID,
		link.Link,
		link.Title,
	)

	// Erro ao criar Link no DB
	if err != nil {
		return fmt.Errorf("erro ao criar link: %w", err)
	}

	return nil
}

// FindLinksFromTrip recupera links com base no ID da viagem
func (repo *LinksRepository) FindLinksFromTrip(tripID string) ([]Links, error) {
	query := `
		SELECT id, trip_id, link, title FROM links WHERE trip_id = $1
	`

	rows, err := repo.DB.Query(query, tripID)
	if err != nil {
		return nil, fmt.Errorf("erro ao encontrar links: %w", err)
	}
	defer rows.Close()

	var links []Links
	for rows.Next() {
		var link Links
		if err := rows.Scan(&link.ID, &link.TripID, &link.Link, &link.Title); err != nil {
			return nil, fmt.Errorf("erro ao escanear link: %w", err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %w", err)
	}

	return links, nil
}
