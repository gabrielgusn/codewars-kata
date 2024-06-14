package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := deck{}.newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := deck{}.newDeck()
	err := d.saveDeckToFile(filename)
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}

func TestDeal(t *testing.T) {
	d := deck{}.newDeck()

	handSize := 5
	hand, remainingDeck := deal(d, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %v, but got %v", handSize, len(hand))
	}

	if len(remainingDeck) != (len(d) - handSize) {
		t.Errorf("Expected remaining deck size of %v, but got %v", len(d)-handSize, len(remainingDeck))
	}
}

func TestShuffle(t *testing.T) {
	d := deck{}.newDeck()
	originalDeck := make(deck, len(d))
	copy(originalDeck, d)

	d.shuffle()

	if reflect.DeepEqual(d, originalDeck) {
		t.Errorf("Expected shuffled deck to be different from original deck")
	}
}
