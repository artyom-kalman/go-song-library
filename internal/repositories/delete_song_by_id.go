package repositories

import (
	"context"
	"fmt"
)

func (repo *SongRepo) DeleteSongById(songId int, ctx context.Context) error {
	query := fmt.Sprintf("DELETE FROM songs WHERE id = %d;", songId)
	rowsAffected, err := repo.conn.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrSongNotFound
	}

	return nil
}
