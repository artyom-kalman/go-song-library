package repositories

import (
	"testing"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
)

func TestAddLyrics(t *testing.T) {
	lyrics := models.NewLyrics{
		SongId: 6,
		Text: []string{
			"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
			"Ooh\nYou set my soul alight\nOoh\nYou set my soul alight\n",
		},
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	err = songRepo.AddLyrycs(&lyrics)
	if err != nil {
		t.Errorf("expected no error\ngot %v", err)
	}

}

func TestAddLyricsInvalidSongId(t *testing.T) {
	lyrics := models.NewLyrics{
		SongId: 1000,
		Text: []string{
			"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
			"Ooh\nYou set my soul alight\nOoh\nYou set my soul alight\n",
		},
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	err = songRepo.AddLyrycs(&lyrics)
	if err == nil {
		t.Errorf("expected no error\ngot %v", err)
	}

}
