package repositories

import "fmt"

func (repo *SongRepo) DeleteSongById(songId int) error {
	query := fmt.Sprintf("DELETE FROM songs WHERE id = %d;", songId)
	rowsAffected, err := repo.conn.Exec(query)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("error deleting a record\n")
	}

	return nil
}
