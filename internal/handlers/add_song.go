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

// @Summary Create a new song
// @Description Add a new song and its lyrics to the database
// @Tags song
// @Accept json
// @Produce json
// @Param song body models.NewSong true "Song information"
// @Success 200 {object} models.Song "Successfully created song"
// @Failure 400 {string} string "Bad request"
// @Failure 405 {string} string "Method not allowed"
// @Failure 500 {string} string "Internal server error"
// @Router /song [post]
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

	song, err := songRepo.AddSong(&newSong)
	if err != nil {
		logger.Error("Failed to add song: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	songLyrics := services.ParseSongText(newSong.Text)
	newLyrics := models.NewLyrics{
		SongId: song.Id,
		Text:   songLyrics,
	}

	err = songRepo.AddLyrycs(&newLyrics)
	if err != nil {
		logger.Error("Failed to add lyrics for the song: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&song); err != nil {
		logger.Error("Failed to encode response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Created new song: %s by %s", newSong.Name, newSong.Group)
}
