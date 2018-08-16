package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Wrong length %v", len(d))
	}
}

// os.remove("_deckTesting")
