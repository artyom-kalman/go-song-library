package models

type Lyrics struct {
	SongId  int    `json:"songId"`
	OrederN int    `json:"orderN"`
	Lyrics  string `json:"lyrics"`
}
