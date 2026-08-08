package engine

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/RewithSolo/arcana-gate/internal/domain"
)

// Config defines execution flags and evaluation strictness for the engine.
type Config struct {
	Strict bool
}

// Engine implements the domain. Oracle interface to provide deterministic fate evaluations.
type Engine struct {
	deck   []domain.Card
	config Config
}

// New creates a new Engine instance initialized with a given deck and configuration.
// If the provided deck is empty, it falls back to the DefaultDeck.
func New(deck []domain.Card, cfg Config) *Engine {
	if len(deck) == 0 {
		deck = DefaultDeck()
	}
	return &Engine{
		deck:   deck,
		config: cfg,
	}
}

// Divine draws a card based on a deterministic seed and evaluates the release status.
func (e *Engine) Divine(ctx context.Context, seed string) (*domain.GateResult, error) {
	if seed == "" {
		return nil, fmt.Errorf("seed cannot be empty for deterministic divination")
	}

	rng := e.createRNG(seed)

	cardIndex := rng.Intn(len(e.deck))
	drawnCard := e.deck[cardIndex]

	position := domain.Upright
	if rng.Float32() < 0.5 {
		position = domain.Reversed
	}

	resultCard := domain.DrawnCard{
		Card:     drawnCard,
		Position: position,
	}

	status := domain.StatusPass
	reason := fmt.Sprintf("Fate allows release with %s (%s)", drawnCard.Name, position)

	if resultCard.IsBlocking() {
		status = domain.StatusBlock
		reason = fmt.Sprintf("Card %s in %s position blocks the pipeline", drawnCard.Name, position)
	} else if position == domain.Reversed && e.config.Strict {
		status = domain.StatusBlock
		reason = fmt.Sprintf("Strict mode enabled: Reversed card %s blocks the pipeline", drawnCard.Name)
	}

	return &domain.GateResult{
		Status:    status,
		DrawnCard: resultCard,
		Reason:    reason,
		Seed:      seed,
	}, nil
}

// createRNG derives a deterministic math/rand source from the SHA-256 hash of the seed.
func (e *Engine) createRNG(seed string) *rand.Rand {
	hash := sha256.Sum256([]byte(seed))
	// Convert the first 8 bytes of the hash into an int64 seed value
	seedInt := int64(binary.BigEndian.Uint64(hash[:8]))
	return rand.New(rand.NewSource(seedInt))
}
