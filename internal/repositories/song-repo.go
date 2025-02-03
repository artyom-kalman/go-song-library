package repositories

import (
	"github.com/artyom-kalman/go-song-library/internal/db"
)

type SongRepo struct {
	conn *db.DBConnection
}

func NewSongRepo(conn *db.DBConnection) *SongRepo {
	return &SongRepo{
		conn: conn,
	}
}
