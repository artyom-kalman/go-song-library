package repositories

import (
	"testing"

	"github.com/artyom-kalman/go-song-library/internal/db"
)

func TestGetSongs(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := SongsSearchParams{
		Offset:   5,
		PageSize: 5,
	}
	_, err = songRepo.GetSongs(&searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
