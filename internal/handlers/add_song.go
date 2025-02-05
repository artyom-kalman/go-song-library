package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func NewSongHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("Received newsong request")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newSong models.NewSong
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		logger.Logger.Error("Error readign request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	err := songRepo.AddSong(&newSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error(fmt.Sprintf("Error adding new song: %s", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	logger.Logger.Info(fmt.Sprintf("Created new song: %s by %s", newSong.Name, newSong.Group))
}
