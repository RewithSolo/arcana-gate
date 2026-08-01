package domain

// Position represents the orientation of a Tarot card.
type Position string

const (
	Upright  Position = "UPRIGHT"
	Reversed Position = "REVERSED"
)

// Suit represents the category or suit of a Tarot card.
type Suit string

const (
	MajorArcana Suit = "MAJOR_ARCANA"
	Wands       Suit = "WANDS"
	Cups        Suit = "CUPS"
	Swords      Suit = "SWORDS"
	Pentacles   Suit = "PENTACLES"
)

// Card defines the structure and attributes of a Tarot card.
type Card struct {
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	Suit          Suit   `json:"suit" yaml:"suit"`
	Value         int    `json:"value" yaml:"value"`
	Description   string `json:"description" yaml:"description"`
	BlockUpright  bool   `json:"block_upright" yaml:"block_upright"`
	BlockReversed bool   `json:"block_reversed" yaml:"block_reversed"`
}

// DrawnCard represents a card drawn in a specific orientation.
type DrawnCard struct {
	Card     Card     `json:"card"`
	Position Position `json:"position"`
}

// IsBlocking determines if the drawn card blocks the pipeline given its orientation.
func (dc DrawnCard) IsBlocking() bool {
	if dc.Position == Upright && dc.Card.BlockUpright {
		return true
	}
	if dc.Position == Reversed && dc.Card.BlockReversed {
		return true
	}
	return false
}
