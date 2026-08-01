package domain

import (
	"testing"
)

func TestDrawnCard_IsBlocking(t *testing.T) {
	tests := []struct {
		name     string
		dc       DrawnCard
		expected bool
	}{
		{
			name: "Upright position and card blocks upright",
			dc: DrawnCard{
				Position: Upright,
				Card: Card{
					BlockUpright: true,
				},
			},
			expected: true,
		},
		{
			name: "Upright position but card does not block upright",
			dc: DrawnCard{
				Position: Upright,
				Card: Card{
					BlockUpright: false,
				},
			},
			expected: false,
		},
		{
			name: "Reversed position and card blocks reversed",
			dc: DrawnCard{
				Position: Reversed,
				Card: Card{
					BlockReversed: true,
				},
			},
			expected: true,
		},
		{
			name: "Reversed position but card does not block reversed",
			dc: DrawnCard{
				Position: Reversed,
				Card: Card{
					BlockReversed: false,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dc.IsBlocking()
			if got != tt.expected {
				t.Errorf("IsBlocking() = %v, want %v", got, tt.expected)
			}
		})
	}
}
