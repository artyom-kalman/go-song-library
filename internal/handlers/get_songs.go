package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetSongsHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("New getsongs request")

	if r.Method != http.MethodGet {
		logger.Error("Wrong HTTP method: %s", r.Method)
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	songRepo := repositories.NewSongRepo(db.GetDatabase())

	searchParams, err := getSongQueryParams(r)
	if err != nil {
		logger.Error("Error parsing query parameters: %v", err)
		http.Error(w, "Error parsing query parameters", http.StatusBadRequest)
		return
	}

	songs, err := songRepo.GetSongs(searchParams)
	if err != nil {
		logger.Error("Error filtering songs: %v", err)
		http.Error(w, "Error filtering songs", http.StatusInternalServerError)
		return
	}

	songsJson, err := json.Marshal(songs)
	if err != nil {
		logger.Error("Error encoding songs to JSON: %v", err)
		http.Error(w, "Error encoding songs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(songsJson)

	logger.Info("Getsongs request served")
}

func getSongQueryParams(r *http.Request) (*repositories.SongQueryParams, error) {
	queryParams := r.URL.Query()

	songId := queryParams.Get("songid")
	songName := queryParams.Get("songname")

	groupId := queryParams.Get("groupid")
	groupName := queryParams.Get("groupname")

	releaseDataStart := queryParams.Get("releasedate-start")
	releaseDateEnd := queryParams.Get("releasedate-end")

	offset := queryParams.Get("offset")
	limit := queryParams.Get("limit")

	searchParams := repositories.NewSongQueryParams()

	searchParams.SongName = songName
	searchParams.GroupName = groupName

	searchParams.StartDate = releaseDataStart
	searchParams.EndDate = releaseDateEnd

	var err error
	if groupId != "" {
		searchParams.GroupId, err = strconv.Atoi(groupId)
		if err != nil {
			return nil, err
		}
	}

	if songId != "" {
		searchParams.SongId, err = strconv.Atoi(songId)
		if err != nil {
			return nil, err
		}
	}

	if offset != "" {
		searchParams.Offset, err = strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
	}

	if limit != "" {
		searchParams.Limit, err = strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
	}

	return searchParams, nil
}

func validateSearchSongParams(searchParams *repositories.SongQueryParams) bool {
	if searchParams.SongId < 0 {
		return false
	}
	if searchParams.GroupId < 0 {
		return false
	}
	if searchParams.Offset < 0 {
		return false
	}
	if searchParams.Limit < 0 {
		return false
	}
	return true
}
