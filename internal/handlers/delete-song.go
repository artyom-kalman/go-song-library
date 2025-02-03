package handlers

import (
	"encoding/json"
	"log"
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

	var req struct {
		SongId string `json:"songId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	songIdStr := req.SongId
	if songIdStr == "" {
		http.Error(w, "Specify song id", http.StatusBadRequest)
		return
	}

	songId, err := strconv.Atoi(songIdStr)
	if err != nil {
		http.Error(w, "Given value is not integer", http.StatusInternalServerError)
		return
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Error connection to database:", err)
		return
	}

	songRepo := repositories.NewSongRepo(dbConn)

	err = songRepo.DeleteSongById(songId)
	if err != nil {
		http.Error(w, "Error finding song with given id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
