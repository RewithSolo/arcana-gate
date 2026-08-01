package domain

import "context"

// DecisionStatus defines the outcome of a gate check.
type DecisionStatus string

const (
	StatusPass  DecisionStatus = "PASS"
	StatusBlock DecisionStatus = "BLOCK"
)

// GateResult holds the final verdict and metadata of the divination process.
type GateResult struct {
	Status    DecisionStatus `json:"status"`
	DrawnCard DrawnCard      `json:"drawn_card"`
	Reason    string         `json:"reason"`
	Seed      string         `json:"seed"`
}

// Oracle defines the interface for executing fate-based gate decisions.
type Oracle interface {
	Divine(ctx context.Context, seed string) (*GateResult, error)
}
