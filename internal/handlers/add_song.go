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
// @Param song body models.NewSongRequest true "Song information"
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

	var newSongRequest *models.NewSongRequest
	if err := json.NewDecoder(r.Body).Decode(&newSongRequest); err != nil {
		logger.Error("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song, err := services.GetSongInfo(newSongRequest)
	if err != nil {
		logger.Error("Failed to get song info: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	var group *models.Group
	if !songRepo.IsGroupExist(song.Group) {
		group, err = songRepo.AddGroup(song.Group)
		if err != nil {
			return
		}
	} else {
		group, err = songRepo.GetGroudByName(song.Group)
		if err != nil {
			return
		}
	}

	newSong, err := songRepo.AddSong(song, group)
	if err != nil {
		logger.Error("Failed to add song: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	songLyrics := services.ParseSongText(song.Text)
	newLyrics := models.NewLyrics{
		SongId: newSong.Id,
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

	logger.Info("Created new song: %s", newSong.Name)
}
