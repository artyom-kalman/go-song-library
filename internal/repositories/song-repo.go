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
	group, err := repo.GetGroudByName(song.Group)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO songs (name, group_id) VALUES ('%s', %d);", song.Name, group.Id)
	_, err = repo.conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (repo *SongRepo) GetGroudByName(groupname string) (*models.Group, error) {
	query := fmt.Sprintf("SELECT id FROM groups WHERE name = '%s';", groupname)
	queryResult, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}

	if !queryResult.Next() {
		return nil, fmt.Errorf("no group with name %s", groupname)
	}

	var groupId int
	err = queryResult.Scan(&groupId)
	if err != nil {
		return nil, fmt.Errorf("error finding group: %v", err)
	}

	return &models.Group{
		Id:   groupId,
		Name: groupname,
	}, nil
}

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
