package engine

import "github.com/RewithSolo/arcana-gate/internal/domain"

// DefaultDeck returns a curated set of Major Arcana cards adapted for IT and DevOps scenarios.
func DefaultDeck() []domain.Card {
	return []domain.Card{
		{
			ID:            "tower",
			Name:          "The Tower",
			Suit:          domain.MajorArcana,
			Value:         16,
			Description:   "Catastrophic infrastructure outage, production crash, and data loss.",
			BlockUpright:  true,
			BlockReversed: true,
		},
		{
			ID:            "death",
			Name:          "Death",
			Suit:          domain.MajorArcana,
			Value:         13,
			Description:   "Irreversible architecture transformation. Danger: potential database wipe.",
			BlockUpright:  true,
			BlockReversed: false,
		},
		{
			ID:            "devil",
			Name:          "The Devil",
			Suit:          domain.MajorArcana,
			Value:         15,
			Description:   "Critical technical debt and unmaintainable code hacks have breached production.",
			BlockUpright:  true,
			BlockReversed: true,
		},
		{
			ID:            "sun",
			Name:          "The Sun",
			Suit:          domain.MajorArcana,
			Value:         19,
			Description:   "Flawless release, 100% test coverage, and absolute uptime stability.",
			BlockUpright:  false,
			BlockReversed: false,
		},
		{
			ID:            "world",
			Name:          "The World",
			Suit:          domain.MajorArcana,
			Value:         21,
			Description:   "Perfect microservices synergy, zero latency, and pure green telemetry.",
			BlockUpright:  false,
			BlockReversed: false,
		},
		{
			ID:            "fool",
			Name:          "The Fool",
			Suit:          domain.MajorArcana,
			Value:         0,
			Description:   "Untested Friday evening release. Pure luck-driven deployment.",
			BlockUpright:  false,
			BlockReversed: true,
		},
		{
			ID:            "high_priestess",
			Name:          "The High Priestess",
			Suit:          domain.MajorArcana,
			Value:         2,
			Description:   "Hidden edge-case bugs, subtle race conditions, and unhandled memory leaks.",
			BlockUpright:  true,
			BlockReversed: false,
		},
		{
			ID:            "wheel_of_fortune",
			Name:          "The Wheel of Fortune",
			Suit:          domain.MajorArcana,
			Value:         10,
			Description:   "Pipeline outcome depends entirely on pure chance and external dependencies.",
			BlockUpright:  false,
			BlockReversed: false,
		},
		{
			ID:            "hanged_man",
			Name:          "The Hanged Man",
			Suit:          domain.MajorArcana,
			Value:         12,
			Description:   "Deadlocks, hanging background workers, and infinite CI/CD build loops.",
			BlockUpright:  true,
			BlockReversed: false,
		},
	}
}
