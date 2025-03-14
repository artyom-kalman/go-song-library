package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// @Summary Delete song by ID
// @Description Delete a song by its ID from the database.
// @Tags song
// @Param songid query integer true "Song ID to delete"
// @Success 200 {string} string "Successfully deleted song"
// @Failure 400 {string} string "Invalid song id"
// @Failure 404 {string} string "Song not found"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /song [delete]
func HandleDeleteSongRequest(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request for deleting a song")

	songId, err := getSongIdFromRequest(r)
	if err != nil {
		logger.Error("Invalid song id: %v", err)
		http.Error(w, "Invalid song id", http.StatusBadRequest)
		return
	}

	err = deleteSongById(songId)
	if err != nil {
		logger.Error("Error deleting song %d: %v", songId, err)
		if err == repositories.ErrSongNotFound {
			http.Error(w, "Song not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	logger.Info("Successfully deleted song id = %d", songId)
}

func getSongIdFromRequest(r *http.Request) (int, error) {
	songIdParam := r.URL.Query().Get("songid")
	if songIdParam == "" {
		return 0, errors.New("Song id is required")
	}

	songId, err := strconv.Atoi(songIdParam)
	if err != nil {
		return 0, errors.New("Song id must be a integer")
	}

	return songId, nil
}

func deleteSongById(songId int) error {
	songRepo := repositories.NewSongRepo(db.GetDatabase())

	return songRepo.DeleteSongById(songId)
}
