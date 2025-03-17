package repositories

import (
	"github.com/artyom-kalman/go-song-library/internal/db"
)

type SongRepo struct {
	conn db.DatabaseConnection
}

func NewSongRepo(conn db.DatabaseConnection) *SongRepo {
	return &SongRepo{
		conn: conn,
	}
}
