package main

import (
	"fmt"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/handlers"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// TODO

// Получение теста песни
//     GET /lyrics?song-id= &offset= &size=

// Получение всей библиотеки
//     GET /songs
//
// Все поля при добавлении песни

const PORT = ":3030"

func main() {
	logger.InitLogger()

	_, err := db.ConnectToDB()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error connection to database: %s", err.Error()))
		return
	}

	http.HandleFunc("/song", handlers.SongHandler)
	http.HandleFunc("/songs", handlers.GetSongsHandler)

	logger.Logger.Info(fmt.Sprintf("Server is running on %s", PORT))

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
