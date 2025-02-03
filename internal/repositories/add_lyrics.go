package repositories

import (
	"fmt"
	"strings"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) AddLyrycs(lyrics *models.NewLyrics) error {
	for oreder, verse := range lyrics.Text {
		verse = strings.ReplaceAll(verse, "'", "''")
		query := fmt.Sprintf("INSERT INTO lyrics (song_id, order_n, lyrics) VALUES (%d, %d, '%s');", lyrics.SongId, oreder+1, verse)

		_, err := repo.conn.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
