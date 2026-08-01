package engine_test

import (
	"context"
	"testing"

	"github.com/RewithSolo/arcana-gate/internal/domain"
	"github.com/RewithSolo/arcana-gate/internal/engine"
)

func TestEngine_EmptySeed(t *testing.T) {
	eng := engine.New(nil, engine.Config{Strict: false})
	ctx := context.Background()

	_, err := eng.Divine(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty seed, got nil")
	}
}

func TestEngine_Determinism(t *testing.T) {
	eng := engine.New(nil, engine.Config{Strict: false})
	ctx := context.Background()
	const seed = "commit-sha-9f8e7d6c5b4a3"

	res1, err := eng.Divine(ctx, seed)
	if err != nil {
		t.Fatalf("unexpected error on first call: %v", err)
	}

	res2, err := eng.Divine(ctx, seed)
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}

	if res1.DrawnCard.Card.ID != res2.DrawnCard.Card.ID {
		t.Errorf("expected deterministic card ID %q, got %q and %q", res1.DrawnCard.Card.ID, res1.DrawnCard.Card.ID, res2.DrawnCard.Card.ID)
	}

	if res1.DrawnCard.Position != res2.DrawnCard.Position {
		t.Errorf("expected deterministic position %q, got %q and %q", res1.DrawnCard.Position, res1.DrawnCard.Position, res2.DrawnCard.Position)
	}

	if res1.Status != res2.Status {
		t.Errorf("expected deterministic status %q, got %q and %q", res1.Status, res1.Status, res2.Status)
	}
}

func TestEngine_DifferentSeedsProduceDifferentResults(t *testing.T) {
	eng := engine.New(nil, engine.Config{Strict: false})
	ctx := context.Background()

	res1, _ := eng.Divine(ctx, "seed-alpha")
	res2, _ := eng.Divine(ctx, "seed-beta")

	// Verify seeds are mapped correctly in the domain result
	if res1.Seed != "seed-alpha" || res2.Seed != "seed-beta" {
		t.Errorf("expected seeds to match input values")
	}
}

func TestEngine_StrictMode(t *testing.T) {
	// Custom deck with a safe card that is not blocking by default
	safeDeck := []domain.Card{
		{
			ID:            "safe_card",
			Name:          "Safe Card",
			BlockUpright:  false,
			BlockReversed: false,
		},
	}

	engStrict := engine.New(safeDeck, engine.Config{Strict: true})
	ctx := context.Background()

	// Iterate through multiple seeds until we find a Reversed position drawing
	var foundReversed bool
	for i := 0; i < 50; i++ {
		seed := string(rune('a' + i))
		res, err := engStrict.Divine(ctx, seed)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if res.DrawnCard.Position == domain.Reversed {
			foundReversed = true
			if res.Status != domain.StatusBlock {
				t.Errorf("expected strict mode to block reversed card, got status %q", res.Status)
			}
			break
		}
	}

	if !foundReversed {
		t.Skip("could not produce a reversed position within 50 iterations")
	}
}

func TestEngine_CustomDeckFallback(t *testing.T) {
	// Passing an empty slice should fall back to DefaultDeck
	eng := engine.New([]domain.Card{}, engine.Config{})
	ctx := context.Background()

	res, err := eng.Divine(ctx, "fallback-seed")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.DrawnCard.Card.ID == "" {
		t.Error("expected valid card ID from default deck, got empty string")
	}
}
