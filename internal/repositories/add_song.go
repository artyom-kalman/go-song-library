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

	query := getNewSongQuery(song, group)
	_, err = repo.conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func getNewSongQuery(song *models.NewSong, group *models.Group) string {
	return fmt.Sprintf(
		"INSERT INTO songs (name, group_id, release_date, link) VALUES ('%s', %d, '%s', '%s');",
		song.Name,
		group.Id,
		formatDate(song.ReleaseDate),
		song.Link,
	)
}

func formatDate(date string) string {
	return fmt.Sprintf("%s-%s-%s",
		date[6:],
		date[3:5],
		date[0:2],
	)
}
