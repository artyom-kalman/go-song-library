package main

import (
	"log"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/handlers"
)

const PORT = ":3030"

func main() {
	_, err := db.ConnectToDB()
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}

	http.HandleFunc("/song", handlers.SongHandler)

	log.Println("Started a server on", PORT)

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln("error starting a server:", err)
	}
}
