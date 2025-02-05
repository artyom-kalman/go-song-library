package repositories

import (
	"testing"

	"github.com/artyom-kalman/go-song-library/internal/db"
)

func TestGetSongsNotReturnError(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()

	_, err = songRepo.GetSongs(searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetSongsByGroupName(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()
	searchParams.GroupName = "Pink Floyd"

	songs, err := songRepo.GetSongs(searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if songs[0].GroupName != "Pink Floyd" {
		t.Errorf("GroupName)Expected Pink Floyd, got %v", songs[0])
	}
}

func TestGetSongsBySongId(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()
	searchParams.SongId = 1

	songs, err := songRepo.GetSongs(searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if songs[0].Id != 1 {
		t.Errorf("Expected song ID 1, got %v", songs[0].Id)
	}
}

func TestGetSongsByGroupId(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()
	searchParams.GroupId = 1

	songs, err := songRepo.GetSongs(searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if songs[0].GroupId != 1 {
		t.Errorf("Expected group ID 1, got %v", songs[0].GroupId)
	}
}

func TestGetSongsByDateRange(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()
	searchParams.StartDate = "2000-01-01"
	searchParams.EndDate = "2023-12-31"

	songs, err := songRepo.GetSongs(searchParams)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(songs) == 0 {
		t.Error("Expected songs in date range, got none")
	}
}

func TestGetSongsWithInvalidDateRange(t *testing.T) {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := NewSongRepo(dbConn)

	searchParams := NewSongQueryParams()
	searchParams.StartDate = "invalid-date"
	searchParams.EndDate = "2023-12-31"

	_, err = songRepo.GetSongs(searchParams)

	if err == nil {
		t.Error("Expected error for invalid start date, got none")
	}
}
