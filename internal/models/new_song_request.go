package models

type NewSongRequest struct {
	Song  string `json:"song"`
	Group string `json:"group"`
}
