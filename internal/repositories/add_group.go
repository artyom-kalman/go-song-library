package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (songRepo *SongRepo) AddGroup(name string) (*models.Group, error) {
	query := fmt.Sprintf("INSERT INTO groups (name) VALUES ('%s') RETURNING id, name", name)

	queryResult, err := songRepo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var newGroup models.Group
	queryResult.Next()
	queryResult.Scan(&newGroup.Id, &newGroup.Name)

	return &newGroup, nil
}
