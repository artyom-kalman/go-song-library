package models

type NewSong struct {
	Name        string `json:"song"`
	Group       string `json:"group"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
}
