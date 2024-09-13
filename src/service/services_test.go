package service

import (
	"testing"
)

func TestGetFix(t *testing.T) {
	newNote := "постовить чяйник"
	fixedWords, _ := GetFix(newNote)
	sug1 := fixedWords[0]
	expected := "поставить"
	if sug1.Sug[0] != expected {
		t.Errorf("Expected %s, got %v", expected, sug1.Sug[0])
	}

	sug2 := fixedWords[1]
	expected = "чайник"
	if sug2.Sug[0] != expected {
		t.Errorf("Expected %s, got %v", expected, sug2.Sug[0])
	}
}

func TestFixNote(t *testing.T) {
	newNote := "постовить чяйник"
	fixedWords, _ := GetFix(newNote)
	fixedText := FixNote(newNote, fixedWords)
	expected := "поставить чайник"
	if fixedText != expected {
		t.Errorf("Expected %s, got %v", expected, fixedText)
	}
}
