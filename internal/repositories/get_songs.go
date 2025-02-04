package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

type SongsSearchParams struct {
	Id          int
	Name        string
	GroupId     int
	GroupName   string
	ReleaseData string
	Offset      int
	PageSize    int
}

func (repo *SongRepo) GetSongs(searchParams *SongsSearchParams) ([]*models.Song, error) {
	query := makeQueryForSongs(searchParams)

	println(query)

	queryResult, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	songs := make([]*models.Song, 0)
	for queryResult.Next() {
		var song models.Song
		err = queryResult.Scan(
			&song.Id,
			&song.Name,
			&song.GroupId,
			&song.GroupName,
			&song.ReleaseDate,
			&song.Link,
		)
		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)
	}

	if err = queryResult.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func makeQueryForSongs(searchParams *SongsSearchParams) string {
	query := fmt.Sprintf(
		`SELECT s.id, s.name AS song_name, g.id AS group_id, g.name AS group_name, s.release_date, s.link
		FROM songs AS s
		INNER JOIN groups AS g ON s.group_id = g.id
		WHERE s.name ILIKE '%%%s%%'
		OFFSET %d
		LIMIT %d`,
		searchParams.Name, searchParams.Offset, searchParams.PageSize,
	)

	return query
}
