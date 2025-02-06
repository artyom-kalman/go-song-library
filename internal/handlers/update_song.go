package handlers

import (
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
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /song [patch]
func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("New updatesong request received")

	if r.Method != http.MethodPatch {
		logger.Info("Invalid request method received: %s", r.Method)
		http.Error(w, "Wrong method: expected PATCH", http.StatusMethodNotAllowed)
		return
	}

	var song models.UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		logger.Info("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	updatedSong, err := songRepo.UpdateSong(&song)
	if err != nil {
		logger.Info("Failed to update song: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&updatedSong); err != nil {
		logger.Info("Failed to encode response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Song update request completed successfully")
}
