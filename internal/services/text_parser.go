package services

import "strings"

func ParseSongText(text string) []string {
	return strings.Split(text, "\n\n")
}
