package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//Code to make sure that a deck is created with x number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	//Code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected firsst card of Ace of Spades, but got %v", d[0])
	}
	//Code to make sure that the last card is a Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected firsst card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

//a test to make sure I cover an edge case that might make a little problem:
// During testing: create a deck -> save to file -> file created -> attempt to load -> CRASH!!
// then after it crashes there would be no place for us doing clean ups
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Delete any files in the current working directory with the name "_decktesting"
	os.Remove("_decktesting")
	//create a new deck
	d := newDeck()
	//save to a file
	d.saveToFile("_decktesting")
	//load from file
	loadedDeck := newDeckFromFile("_decktesting")
	//assert deck length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	//delete any files in current working directory with the name "_decktesting"
	os.Remove("_decktesting")
}
