package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetSongsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("New getsongs request")

	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	songRepo := repositories.NewSongRepo(dbConn)

	searchParams := repositories.SongsSearchParams{}

	songs, err := songRepo.GetSongs(&searchParams)
	if err != nil {
		logger.Logger.Debug(err.Error())
		http.Error(w, "Error filtering songs", http.StatusInternalServerError)
		return
	}

	songsJson, err := json.Marshal(songs)
	if err != nil {
		logger.Logger.Debug(err.Error())
		http.Error(w, "Error encoding songs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJson)
}
