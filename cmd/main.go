package main

import (
	"fmt"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/handlers"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// TODO

// Все поля при добавлении песни

const PORT = ":3030"

func main() {
	logger.InitLogger()

	databaseConfig, err := config.LoadDBConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	err = db.InitDatabase(databaseConfig)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	http.HandleFunc("/song", handlers.SongHandler)
	http.HandleFunc("/songs", handlers.GetSongsHandler)
	http.HandleFunc("/lyrics", handlers.GetLyricsHandler)

	logger.Logger.Info(fmt.Sprintf("Server is running on %s", PORT))

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
