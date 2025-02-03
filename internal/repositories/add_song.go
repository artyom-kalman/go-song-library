package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) AddSong(song *models.NewSong) error {
	group, err := repo.GetGroudByName(song.Group)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO songs (name, group_id) VALUES ('%s', %d);", song.Name, group.Id)
	_, err = repo.conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}
