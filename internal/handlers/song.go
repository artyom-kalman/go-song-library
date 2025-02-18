package handlers

import "net/http"

func SongHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddSongHandler(w, r)
	case http.MethodPatch:
		UpdateSongHandler(w, r)
	case http.MethodDelete:
		DeleteSongHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
