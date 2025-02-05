package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/internal/services"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func AddSongHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received newsong request")

	if r.Method != http.MethodPost {
		logger.Error("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newSong models.NewSong
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		logger.Error("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.GetSongInfo(&newSong); err != nil {
		logger.Error("Failed to get song info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	if err := songRepo.AddSong(&newSong); err != nil {
		logger.Error("Failed to add song: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	logger.Info("Created new song: %s by %s", newSong.Name, newSong.Group)
}
