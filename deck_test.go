package main

import (
	"os"
	"testing"
)

// new Deck should have 4 items
//t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected Deck length of 16 but got %v", len(d))
	}
	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected Ace of Diamonds to be the first card but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs to be the last card but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.savetoFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected to load 16 from file but got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
