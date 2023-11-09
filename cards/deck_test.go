package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last element to be Kings of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeskFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	deckFromFile := newDeckFromFile("_decktesting")

	if len(deck) != len(deckFromFile) {
		t.Errorf("The expected length was %v but got %v", len(deck), len(deckFromFile))
	}

	os.Remove("_decktesting")
}
