package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"my-app/internal/pkg/models"
	"time"
)

func InsertTest(test models.Test) (string, error) {
	query := `INSERT INTO test (id, name, created_at, updated_at)
              VALUES ($1, $2, $3, $4) RETURNING id`
	var returnedId string
	err := DB.QueryRow(context.Background(), query,
		uuid.New(),
		test.Name,
		time.Now().UTC(),
		time.Now().UTC()).Scan(&returnedId)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return returnedId, nil
}

func DeleteTestById(testId string) error {
	query := `DELETE FROM test WHERE id = $1`

	// Execute the DELETE query
	_, err := DB.Exec(context.Background(), query, testId)
	if err != nil {
		return fmt.Errorf("failed to delete test with id %s: %w", testId, err)
	}

	return nil
}

func GetTestIdsByName(name string) ([]string, error) {
	query := `SELECT id FROM test WHERE name = $1`
	rows, err := DB.Query(context.Background(), query, name)

	if err != nil {
		log.Fatal("Error fetching ids for name", err)
	}
	defer rows.Close()

	// Collect IDs into a slice
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Fatal("Failed to scan row: ", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Row iteration error: ", err)
	}

	return ids, err
}
