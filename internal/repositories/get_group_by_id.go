package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) GetGroudByName(groupname string) (*models.Group, error) {
	query := fmt.Sprintf("SELECT id FROM groups WHERE name = '%s';", groupname)
	queryResult, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	if !queryResult.Next() {
		return nil, fmt.Errorf("no group with name %s", groupname)
	}

	var groupId int
	err = queryResult.Scan(&groupId)
	if err != nil {
		return nil, fmt.Errorf("error finding group: %v", err)
	}

	return &models.Group{
		Id:   groupId,
		Name: groupname,
	}, nil
}
