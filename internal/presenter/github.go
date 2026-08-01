package presenter

import (
	"fmt"
	"os"

	"github.com/RewithSolo/arcana-gate/internal/domain"
)

// GitHubPresenter formats gate outcomes as Markdown for GitHub Actions Step Summaries.
type GitHubPresenter struct {
	summaryPath string
}

// NewGitHubPresenter creates a presenter for writing GitHub Step Summaries.
// Uses the $GITHUB_STEP_SUMMARY environment variable path.
func NewGitHubPresenter() *GitHubPresenter {
	return &GitHubPresenter{
		summaryPath: os.Getenv("GITHUB_STEP_SUMMARY"),
	}
}

// Render appends a formatted Markdown card to the GitHub Step Summary file if available.
func (g *GitHubPresenter) Render(res *domain.GateResult) error {
	if g.summaryPath == "" {
		// Not running inside a GitHub Actions environment that supports step summary
		return nil
	}
	statusEmoji := "✅"
	if res.Status == domain.StatusBlock {
		statusEmoji = "❌"
	}
	markdown := fmt.Sprintf(`## %s Arcana Gate Decision

| Property | Value |
| :--- | :--- |
| **Drawn Card** | %s |
| **Position** | %s |
| **Status** | %s **%s** |
| **Determinism Seed** | `+"`%s`"+` |

> **Divination Insight:**
> %s
*%s*
`,
		statusEmoji,
		res.DrawnCard.Card.Name,
		res.DrawnCard.Position,
		statusEmoji,
		res.Status,
		res.Seed,
		res.DrawnCard.Card.Description,
		res.Reason,
	)
	file, err := os.OpenFile(g.summaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open GITHUB_STEP_SUMMARY file: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(markdown); err != nil {
		return fmt.Errorf("failed to write to GITHUB_STEP_SUMMARY: %w", err)
	}
	return nil
}
