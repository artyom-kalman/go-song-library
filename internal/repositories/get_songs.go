package repositories

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) GetSongs(queryParams *SongQueryParams, ctx context.Context) ([]*models.Song, error) {
	query := makeQueryForGetSongs(queryParams)

	queryResult, err := repo.conn.QueryContext(ctx, query)
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

func makeQueryForGetSongs(queryParams *SongQueryParams) string {
	query := fmt.Sprintf(`SELECT s.id, s.name AS song_name, g.id AS group_id, g.name AS group_name, s.release_date, s.link
		FROM songs AS s
		INNER JOIN groups AS g ON s.group_id = g.id WHERE s.name ILIKE '%%%s%%'`, queryParams.SongName)

	if queryParams.StartDate != "" {
		query = fmt.Sprintf("%s AND s.release_date >= '%s'", query, queryParams.StartDate)
	}

	if queryParams.EndDate != "" {
		query = fmt.Sprintf("%s AND s.release_date < '%s'", query, queryParams.EndDate)
	}

	if queryParams.GroupName != "" {
		query = fmt.Sprintf("%s AND g.name ILIKE '%%%s%%'", query, queryParams.GroupName)
	}

	if queryParams.GroupId >= 0 {
		query = fmt.Sprintf("%s AND g.id = %d", query, queryParams.GroupId)
	}

	if queryParams.SongId >= 0 {
		query = fmt.Sprintf("%s AND s.id = %d", query, queryParams.SongId)
	}

	if queryParams.Offset >= 0 {
		query = fmt.Sprintf("%s OFFSET %d", query, queryParams.Offset)
	}

	if queryParams.Limit >= 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, queryParams.Limit)
	}

	return query
}
