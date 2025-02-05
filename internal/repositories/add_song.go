package repositories

import (
	"database/sql"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) AddSong(song *models.NewSong) (int, error) {
	group, err := repo.GetGroudByName(song.Group)
	if err != nil {
		return 0, err
	}

	query := getNewSongQuery(song, group)
	queryResult, err := repo.conn.Query(query)
	if err != nil {
		return 0, err
	}

	songId := getNewSongIdFromQueryResult(queryResult)

	return songId, nil
}

func getNewSongIdFromQueryResult(queryResult *sql.Rows) int {
	var songId int

	queryResult.Next()
	queryResult.Scan(&songId)

	return songId
}

func getNewSongQuery(song *models.NewSong, group *models.Group) string {
	return fmt.Sprintf(
		"INSERT INTO songs (name, group_id, release_date, link) VALUES ('%s', %d, '%s', '%s') RETURNING id;",
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
