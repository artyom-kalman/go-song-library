package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetLyricsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("Received GetLyrics request")

	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	db, err := db.ConnectToDB()
	if err != nil {
		logger.Logger.Debug("Error connection to database")
		http.Error(w, "Error processign request", http.StatusInternalServerError)
		return
	}

	repo := repositories.NewSongRepo(db)

	queryParams, err := getLyricsQueryParamsFromRequest(r)
	if err != nil {
		logger.Logger.Debug(err.Error())
		http.Error(w, "Invalid arguments", http.StatusBadRequest)
		return
	}

	lyrics, err := repo.GetLyrics(queryParams)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(lyrics)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	logger.Logger.Info("Successfully processed GetLyrics request")
}

func getLyricsQueryParamsFromRequest(r *http.Request) (*repositories.LyricsQueryParams, error) {
	requestParams := r.URL.Query()

	result := repositories.LyricsQueryParams{}

	if requestParams.Has("songid") {
		songId, err := strconv.Atoi(requestParams.Get("songid"))
		if err != nil {
			return nil, err
		}

		result.SongId = songId
	} else {
		return nil, fmt.Errorf("Provide song id")
	}

	if requestParams.Has("offset") {
		offset, err := strconv.Atoi(requestParams.Get("offset"))
		if err != nil {
			return nil, err
		}

		result.Offset = offset
	}

	if requestParams.Has("limit") {
		limit, err := strconv.Atoi(requestParams.Get("limit"))
		if err != nil {
			return nil, err
		}

		result.Limit = limit
	}

	return &result, nil
}
