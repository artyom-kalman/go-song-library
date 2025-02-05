package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

type LyricsQueryParams struct {
	SongId int
	Offset int
	Limit  int
}

func (repo *SongRepo) GetLyrics(queryParams *LyricsQueryParams) ([]*models.Lyrics, error) {
	query := makeGetLyricsQuery(queryParams)

	rows, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	lyricsArray := make([]*models.Lyrics, 0)
	for rows.Next() {
		var lyrics models.Lyrics
		err = rows.Scan(
			&lyrics.SongId,
			&lyrics.OrederN,
			&lyrics.Lyrics,
		)
		if err != nil {
			return nil, err
		}

		lyricsArray = append(lyricsArray, &lyrics)
	}

	return lyricsArray, nil
}

func makeGetLyricsQuery(queryParams *LyricsQueryParams) string {
	query := fmt.Sprintf("SELECT song_id, order_n, lyrics FROM lyrics WHERE song_id = %d", queryParams.SongId)

	if queryParams.Offset > 0 {
		query = fmt.Sprintf("%s AND order_n > %d", query, queryParams.Offset)
	}

	if queryParams.Limit > 0 && queryParams.Offset > 0 {
		limit := queryParams.Offset + queryParams.Limit
		query = fmt.Sprintf("%s AND order_n <= %d", query, limit)
	}

	return query
}
