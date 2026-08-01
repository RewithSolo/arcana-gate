package domain

import (
	"testing"
)

func TestGateResult(t *testing.T) {
	res := GateResult{
		Status: "PASS",
		Reason: "Divine alignment complete",
		Seed:   "commit-sha-123",
		DrawnCard: DrawnCard{
			Card: Card{
				Name:        "The Star",
				Description: "Hope and inspiration",
			},
			Position: "upright",
		},
	}

	if res.Status != "PASS" {
		t.Errorf("expected status 'PASS', got %s", res.Status)
	}

	if res.DrawnCard.Card.Name != "The Star" {
		t.Errorf("expected card name 'The Star', got %s", res.DrawnCard.Card.Name)
	}
}
