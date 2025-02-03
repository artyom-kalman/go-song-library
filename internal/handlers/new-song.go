package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
)

func NewSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newSong models.NewSong
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("error connection to database:", err)
		return
	}

	songRepo := repositories.NewSongRepo(dbConn)

	err = songRepo.NewSong(&newSong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("error adding new song", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Printf("Created new song: %s by %s\n", newSong.Name, newSong.Group)
}
