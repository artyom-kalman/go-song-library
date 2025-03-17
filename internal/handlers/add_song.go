package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/internal/services"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// @Summary Create a new song
// @Description Add a new song and its lyrics to the database
// @Tags song
// @Accept json
// @Produce json
// @Param song body models.NewSongRequest true "Song information"
// @Success 200 {object} models.Song
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /song [post]
func HandleAddSongRequest(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request to add a new song")

	var newSongRequest models.NewSongRequest
	if err := json.NewDecoder(r.Body).Decode(&newSongRequest); err != nil {
		logger.Error("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newSong, err := addSong(&newSongRequest, r.Context())
	if err != nil {
		logger.Error("Error adding new song: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&newSong); err != nil {
		logger.Error("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created song: %s", newSong.Name)
}

func addSong(newSongRequest *models.NewSongRequest, ctx context.Context) (*models.Song, error) {
	logger.Debug("New song request body: %+v", newSongRequest)

	songRepo := repositories.NewSongRepo(db.Database())

	song, err := services.GetSongInfo(newSongRequest)
	if err != nil {
		return nil, err
	}

	group, err := findGroup(song.Group, ctx)
	if err != nil {
		return nil, err
	}

	newSong, err := songRepo.AddSong(song, group, ctx)
	if err != nil {
		return nil, err
	}

	songLyrics := services.ParseSongText(song.Text)
	newLyrics := models.NewLyrics{
		SongId: newSong.Id,
		Text:   songLyrics,
	}

	err = songRepo.AddLyrics(&newLyrics, ctx)
	if err != nil {
		return nil, err
	}

	return newSong, nil
}

func findGroup(groupName string, ctx context.Context) (*models.Group, error) {
	songRepo := repositories.NewSongRepo(db.Database())

	if songRepo.IsGroupExist(groupName, ctx) {
		g, err := songRepo.GetGroudByName(groupName, ctx)
		if err != nil {
			return nil, err
		}
		return g, nil
	}

	g, err := songRepo.AddGroup(groupName)
	if err != nil {
		return nil, err
	}
	return g, nil
}
