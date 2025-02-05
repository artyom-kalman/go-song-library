package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received deletesong request")

	if r.Method != http.MethodDelete {
		logger.Error("Invalid request method: %s", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	songId, err := getSongIdFromRequest(r)
	if err != nil {
		logger.Error("Failed to get song ID from request: %v", err)
		http.Error(w, "Require song id", http.StatusBadRequest)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	err = songRepo.DeleteSongById(songId)
	if err != nil {
		logger.Error("Failed to delete song %d: %v", songId, err)
		http.Error(w, "Error finding song with given id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	logger.Info("Deleted song %d", songId)

}

func getSongIdFromRequest(r *http.Request) (int, error) {
	if !r.URL.Query().Has("songid") {
		return 0, fmt.Errorf("Provide songid")
	}

	songId, err := strconv.Atoi(r.URL.Query().Get("songid"))
	if err != nil {
		return 0, nil
	}

	return songId, nil
}
