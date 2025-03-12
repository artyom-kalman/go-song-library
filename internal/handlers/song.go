package handlers

import "net/http"

func HandleSongRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		HandleAddSongRequest(w, r)
	case http.MethodPatch:
		HandleUpdateSongRequest(w, r)
	case http.MethodDelete:
		HandleDeleteSongRequest(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
