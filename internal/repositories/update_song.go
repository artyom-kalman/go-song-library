package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (songRepo *SongRepo) UpdateSong(song *models.UpdateSongRequestBody) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET name = '%s', release_date = '%s', link = '%s' WHERE id = %d RETURNING id, name, group_id, release_date, link",
		song.Name, song.ReleaseDate, song.Link, song.Id,
	)

	queryResult, err := songRepo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	isSongUpdated := queryResult.Next()
	if !isSongUpdated {
		return nil, ErrSongNotFound
	}

	var updatedSong models.Song
	queryResult.Scan(
		&updatedSong.Id,
		&updatedSong.Name,
		&updatedSong.GroupId,
		&updatedSong.ReleaseDate,
		&updatedSong.Link,
	)

	return &updatedSong, nil
}
