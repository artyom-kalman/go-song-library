package services

import (
	"reflect"
	"testing"
)

func TestParseSongTextTwoVerses(t *testing.T) {
	text := "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight\n"

	actual := ParseSongText(text)
	expected := []string{
		"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
		"Ooh\nYou set my soul alight\nOoh\nYou set my soul alight\n",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ParseSongText failed: expected %v \ngot %v", expected, actual)
	}
}

func TestParseSongTestOneVerse(t *testing.T) {
	text := "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?"

	actual := ParseSongText(text)
	expected := []string{
		"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ParseSongText failed: expected %v \ngot %v", expected, actual)
	}
}
