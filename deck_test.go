package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Spades of Ace , but got %v", d[0])
	}

	if d[len(d)-1] != "Clover of Four" {
		t.Errorf("Expected last card of Clover of Four, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	deckToCompare := newDeckFromFile("_decktesting")

	if len(deckToCompare) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	os.Remove("_decktesting")
}
