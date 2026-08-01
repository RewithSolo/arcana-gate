package presenter

import (
	"bytes"
	"testing"

	"github.com/RewithSolo/arcana-gate/internal/domain"
)

func TestTerminalPresenter_Render(t *testing.T) {
	tests := []struct {
		name   string
		result domain.GateResult
	}{
		{
			name: "Render Passed Gate",
			result: domain.GateResult{
				Status: "PASS",
				DrawnCard: domain.DrawnCard{
					Card: domain.Card{
						Name:        "The Sun",
						Description: "Success and joy",
					},
					Position: "upright",
				},
				Reason: "Divine check passed",
				Seed:   "test-seed-123",
			},
		},
		{
			name: "Render Blocked Gate",
			result: domain.GateResult{
				Status: "BLOCK",
				DrawnCard: domain.DrawnCard{
					Card: domain.Card{
						Name:        "The Tower",
						Description: "Destruction and chaos",
					},
					Position: "reversed",
				},
				Reason: "Blocked by divine check",
				Seed:   "test-seed-456",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			p := NewTerminalPresenter(&buf)

			err := p.Render(&tt.result)
			if err != nil {
				t.Fatalf("unexpected error during render: %v", err)
			}

			if buf.Len() == 0 {
				t.Errorf("expected non-empty output buffer")
			}
		})
	}
}
