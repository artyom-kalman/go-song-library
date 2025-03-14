package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (songRepo *SongRepo) GetSongById(id int) (*models.Song, error) {
	query := fmt.Sprintf(
		"SELECT id, name, group_id, release_date, link FROM songs WHERE id = %d",
		id,
	)

	queryResult, err := songRepo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	isSongUpdated := queryResult.Next()
	if !isSongUpdated {
		return nil, ErrSongNotFound
	}

	var song models.Song
	queryResult.Scan(
		&song.Id,
		&song.Name,
		&song.GroupId,
		&song.ReleaseDate,
		&song.Link,
	)

	return &song, nil
}
