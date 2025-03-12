package models

type UpdateSongRequestBody struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}
