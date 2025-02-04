package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) GetSongs(searchParams *SongSearchParams) ([]*models.Song, error) {
	query := makeQueryForSongs(searchParams)

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

	return songs, nil
}

func makeQueryForSongs(searchParams *SongSearchParams) string {
	query := fmt.Sprintf(`SELECT s.id, s.name AS song_name, g.id AS group_id, g.name AS group_name, s.release_date, s.link
		FROM songs AS s
		INNER JOIN groups AS g ON s.group_id = g.id WHERE s.name ILIKE '%%%s%%'`, searchParams.SongName)

	if searchParams.StartDate != "" {
		query = fmt.Sprintf("%s AND s.release_date >= '%s'", query, searchParams.StartDate)
	}

	if searchParams.EndDate != "" {
		query = fmt.Sprintf("%s AND s.release_date < '%s'", query, searchParams.EndDate)
	}

	if searchParams.GroupName != "" {
		query = fmt.Sprintf("%s AND g.name ILIKE '%%%s%%'", query, searchParams.GroupName)
	}

	if searchParams.GroupId >= 0 {
		query = fmt.Sprintf("%s AND g.id = %d", query, searchParams.GroupId)
	}

	if searchParams.SongId >= 0 {
		query = fmt.Sprintf("%s AND s.id = %d", query, searchParams.SongId)
	}

	if searchParams.Offset >= 0 {
		query = fmt.Sprintf("%s OFFSET %d", query, searchParams.Offset)
	}

	if searchParams.Limit >= 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, searchParams.Limit)
	}

	return query
}
