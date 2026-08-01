package presenter

import (
	"fmt"
	"io"
	"strings"

	"github.com/RewithSolo/arcana-gate/internal/domain"
)

// ANSI color codes for rich terminal styling.
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

// TerminalPresenter outputs human-readable ASCII representations of gate decisions to the console.
type TerminalPresenter struct {
	writer io.Writer
}

// NewTerminalPresenter initializes a new console output presenter.
func NewTerminalPresenter(w io.Writer) *TerminalPresenter {
	return &TerminalPresenter{writer: w}
}

// Render formats and prints the GateResult with ASCII borders and ANSI color coding.
func (p *TerminalPresenter) Render(res *domain.GateResult) error {
	statusColor := colorGreen
	if res.Status == domain.StatusBlock {
		statusColor = colorRed
	}

	var positionText string
	if res.DrawnCard.Position == domain.Reversed {
		positionText = "REVERSED"
	} else {
		positionText = "UPRIGHT"
	}

	// Precise padding calculation for the status line to guarantee a 40-char inner width
	statusStr := string(res.Status)
	labelPrefix := "  Status: "
	// Inner width: 1 space left + 9 chars labelPrefix + statusStr + padding + 1 space right = 40 chars
	paddingLen := 40 - (1 + len(labelPrefix) + len(statusStr) + 1)
	if paddingLen < 0 {
		paddingLen = 0
	}
	statusPadding := strings.Repeat(" ", paddingLen)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s%s══════════════════════════════════════════════════%s\n", colorBold, colorCyan, colorReset))
	sb.WriteString(fmt.Sprintf("%s          🔮 ARCANA GATE DECISION 🔮          %s\n", colorBold, colorReset))
	sb.WriteString(fmt.Sprintf("%s%s══════════════════════════════════════════════════%s\n", colorBold, colorCyan, colorReset))
	sb.WriteString(" ┌────────────────────────────────────────┐ \n")
	sb.WriteString(fmt.Sprintf(" │ %s │ \n", centerText(res.DrawnCard.Card.Name, 38)))
	sb.WriteString(fmt.Sprintf(" │ %s │ \n", centerText("["+positionText+"]", 38)))
	sb.WriteString(" │                                        │ \n")
	sb.WriteString(fmt.Sprintf(" │%s%s%s%s%s │ \n", labelPrefix, statusColor+colorBold, statusStr, colorReset, statusPadding))
	sb.WriteString(" └────────────────────────────────────────┘ \n")
	sb.WriteString(fmt.Sprintf("%sDetails:%s %s\n", colorYellow, colorReset, res.Reason))
	sb.WriteString(fmt.Sprintf("%sMeaning:%s %s\n", colorCyan, colorReset, res.DrawnCard.Card.Description))
	sb.WriteString(fmt.Sprintf("%sSeed:%s    %s\n", colorBold, colorReset, res.Seed))
	sb.WriteString(fmt.Sprintf("%s══════════════════════════════════════════════════%s\n", colorCyan, colorReset))

	_, err := fmt.Fprint(p.writer, sb.String())
	return err
}

// centerText computes precise rune-based padding to avoid border misalignment caused by UTF-8 characters.
func centerText(s string, width int) string {
	runes := []rune(s)
	runeLen := len(runes)

	if runeLen >= width {
		return string(runes[:width])
	}

	padding := (width - runeLen) / 2
	extra := (width - runeLen) % 2

	return strings.Repeat(" ", padding) + s + strings.Repeat(" ", padding+extra)
}
