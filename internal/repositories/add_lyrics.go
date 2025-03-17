package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (repo *SongRepo) AddLyrics(lyrics *models.NewLyrics, ctx context.Context) error {
	for order, verse := range lyrics.Text {
		verse = strings.ReplaceAll(verse, "'", "''")
		query := fmt.Sprintf("INSERT INTO lyrics (song_id, order_n, lyrics) VALUES (%d, %d, '%s');", lyrics.SongId, order+1, verse)

		_, err := repo.conn.ExecContext(ctx, query)
		if err != nil {
			return err
		}
	}

	return nil
}
