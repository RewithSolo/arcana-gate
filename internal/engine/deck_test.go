package engine

import (
	"testing"
)

func TestDefaultDeck(t *testing.T) {
	deck := DefaultDeck()

	if len(deck) == 0 {
		t.Fatalf("expected non-empty deck, got 0 cards")
	}

	firstCard := deck[0]

	if firstCard.ID == "" {
		t.Errorf("expected card ID to be set")
	}

	if firstCard.Name == "" {
		t.Errorf("expected card Name to be set")
	}

	if firstCard.Description == "" {
		t.Errorf("expected card Description to be set")
	}
}
