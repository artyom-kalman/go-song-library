package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetSongInfo(song *models.NewSongRequest) (*models.NewSong, error) {
	url := getApiUrl(song.Song, song.Group)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting song info")
	}

	var newSong models.NewSong
	if err := json.NewDecoder(resp.Body).Decode(&newSong); err != nil {
		return nil, err
	}
	newSong.Name = song.Song
	newSong.Group = song.Group

	logger.Debug("Added info to song %s %s", song.Song, song.Group)
	return &newSong, nil
}

func getApiUrl(songName string, groupName string) string {
	baseUrl := fmt.Sprintf("%s/info", config.GetSongInfoApi())
	params := url.Values{}
	params.Add("song", songName)
	params.Add("group", groupName)

	return fmt.Sprintf("%s?%s", baseUrl, params.Encode())
}
