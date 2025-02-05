package repositories

import (
	"testing"

	"github.com/artyom-kalman/go-song-library/internal/db"
)

func TestGetLyricsNoError(t *testing.T) {
	repo := NewSongRepo(db.GetDatabase())

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

func TestGetLyricsWithOffsetNoError(t *testing.T) {
	repo := NewSongRepo(db.GetDatabase())

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
	repo := NewSongRepo(db.GetDatabase())

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
