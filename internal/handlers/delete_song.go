package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
)

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	songId, err := getSongIdFromRequest(r)
	if err != nil {
		http.Error(w, "Require song id", http.StatusBadRequest)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	err = songRepo.DeleteSongById(songId)
	if err != nil {
		http.Error(w, "Error finding song with given id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getSongIdFromRequest(r *http.Request) (int, error) {
	var req struct {
		SongId string `json:"songId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return 0, err
	}

	songIdStr := req.SongId
	if songIdStr == "" {
		return 0, fmt.Errorf("song id must not be empty string")
	}

	songId, err := strconv.Atoi(songIdStr)
	if err != nil {
		return 0, nil
	}

	return songId, nil

}
