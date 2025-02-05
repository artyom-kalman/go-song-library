package repositories

import (
	"testing"

	"github.com/artyom-kalman/go-song-library/internal/db"
)

func TestGetLyricsNoError(t *testing.T) {
	db, _ := db.ConnectToDB()

	repo := NewSongRepo(db)

	queryParams := LyricsQueryParams{
		SongId: 6,
		Offset: 0,
		Limit:  0,
	}

	_, err := repo.GetLyrics(&queryParams)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}
func TestGetLyricsNoLinesNoError(t *testing.T) {
	db, _ := db.ConnectToDB()

	repo := NewSongRepo(db)

	queryParams := LyricsQueryParams{
		SongId: 6,
		Offset: 0,
		Limit:  -1,
	}

	lyrics, err := repo.GetLyrics(&queryParams)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if len(lyrics) != 0 {
		t.Errorf("Expected empty lyrics array, got length %d", len(lyrics))
	}
}

func TestGetLyricsWithOffsetNoError(t *testing.T) {
	db, _ := db.ConnectToDB()

	repo := NewSongRepo(db)

	queryParams := LyricsQueryParams{
		SongId: 6,
		Offset: 2,
		Limit:  0,
	}

	_, err := repo.GetLyrics(&queryParams)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}

func TestGetLyricsWithLimitNoError(t *testing.T) {
	db, _ := db.ConnectToDB()

	repo := NewSongRepo(db)

	queryParams := LyricsQueryParams{
		SongId: 6,
		Offset: 0,
		Limit:  2,
	}

	lyrics, err := repo.GetLyrics(&queryParams)

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if len(lyrics) > 2 {
		t.Errorf("Expected max 2 lyrics lines, got %d", len(lyrics))
	}
}
