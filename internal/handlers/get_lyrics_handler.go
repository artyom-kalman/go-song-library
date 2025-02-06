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

// @Summary Get song lyrics
// @Description Returns lyrics for a specific song
// @Tags lyrics
// @Accept json
// @Produce json
// @Param songid query int true "Song ID"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} models.Lyrics
// @Failure 400 {object} string "Invalid arguments"
// @Failure 405 {object} string "Wrong method"
// @Failure 500 {object} string "Error processing request"
// @Router /lyrics [get]
func GetLyricsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("Received GetLyrics request")

	if r.Method != http.MethodGet {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	repo := repositories.NewSongRepo(db.GetDatabase())

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
