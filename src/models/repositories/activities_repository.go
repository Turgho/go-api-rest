package models

import (
	"database/sql"
	"fmt"
)

type Activity struct {
	ID       string `json:"id"`
	TripID   string `json:"-"`
	Title    string `json:"title"`
	OccursAt string `json:"occurs_at"`
}

type ActivitiesRepository struct {
	DB *sql.DB
}

func NewActivitiesRepository(db *sql.DB) *ActivitiesRepository {
	return &ActivitiesRepository{DB: db}
}

func (repo *ActivitiesRepository) CreateActivity(activity *Activity) error {
	fmt.Println(activity)

	query := `
		INSERT INTO activities
			(id, trip_id, title, occurs_at)
		VALUES
			($1, $2, $3, $4)
	`

	_, err := repo.DB.Exec(
		query,
		activity.ID,
		activity.TripID,
		activity.Title,
		activity.OccursAt,
	)

	// Erro ao criar Atividade no DB
	if err != nil {
		return fmt.Errorf("erro ao criar atividade: %w", err)
	}

	return nil
}

func (repo *ActivitiesRepository) FindActivityFromTrip(tripID string) ([]Activity, error) {
	query := `
		SELECT * FROM activities WHERE trip_id = $1
	`

	rows, err := repo.DB.Query(query, tripID)
	if err != nil {
		return nil, fmt.Errorf("erro ao encontrar atividade: %w", err)
	}
	defer rows.Close()

	var activities []Activity
	for rows.Next() {
		var activity Activity
		if err := rows.Scan(&activity.ID, &activity.TripID, &activity.Title, &activity.OccursAt); err != nil {
			return nil, fmt.Errorf("erro ao escanear atividade: %w", err)
		}
		activities = append(activities, activity)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as linhas: %w", err)
	}

	return activities, nil
}
