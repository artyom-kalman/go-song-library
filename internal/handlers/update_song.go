package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// @Summary Update a song
// @Description Update an existing song in the library
// @Tags song
// @Accept json
// @Produce json
// @Param song body models.UpdateSongRequest true "Song update info"
// @Success 200 {object} models.Song
// @Failure 400 {string} string "Error decoding request body"
// @Failure 404 {string} string "Song not found"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /song [patch]
func HandleUpdateSongRequest(w http.ResponseWriter, r *http.Request) {
	logger.Info("Recieved request to update song")

	var song models.UpdateSongRequestBody
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		logger.Info("Error decoding request body: %v", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	logger.Debug("Update song request body: %+v", song)

	updatedSong, err := updateSong(&song, r.Context())
	if err != nil {
		logger.Info("Error updating song: %v", err)
		if err == repositories.ErrSongNotFound {
			http.Error(w, "Song not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	logger.Debug("Updated song: %+v", updatedSong)

	if err := json.NewEncoder(w).Encode(updatedSong); err != nil {
		logger.Info("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated song")
}

func updateSong(song *models.UpdateSongRequestBody, ctx context.Context) (*models.Song, error) {
	songRepo := repositories.NewSongRepo(db.Database())

	updatedSong, err := songRepo.UpdateSong(song, ctx)
	if err != nil {
		return nil, err
	}

	return updatedSong, nil
}
