package repositories

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
)

type SongRepo struct {
	conn *db.DBConnection
}

func NewSongRepo(conn *db.DBConnection) *SongRepo {
	return &SongRepo{
		conn: conn,
	}
}

func (repo *SongRepo) NewSong(song *models.NewSong) error {
	var groupId int

	query := fmt.Sprintf("SELECT id FROM groups WHERE name = '%s';", song.Group)

	queryResult, err := repo.conn.Query(query)
	if err != nil {
		return err
	}

	if !queryResult.Next() {
		return fmt.Errorf("no group with name %s", song.Group)
	}

	err = queryResult.Scan(&groupId)

	if err != nil {
		return fmt.Errorf("error finding group: %v", err)
	}

	query = fmt.Sprintf("INSERT INTO songs (name, group_id) VALUES ('%s', %d);", song.Name, groupId)
	_, err = repo.conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}
