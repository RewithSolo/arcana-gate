package presenter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/RewithSolo/arcana-gate/internal/domain"
)

func TestGitHubPresenter_Render(t *testing.T) {
	sampleResult := domain.GateResult{
		Status: "PASS",
		DrawnCard: domain.DrawnCard{
			Card: domain.Card{
				Name:        "Wheel of Fortune",
				Description: "Destiny turns",
			},
			Position: "upright",
		},
		Reason: "Good karma",
		Seed:   "test-seed-789",
	}

	t.Run("Success Render to Step Summary", func(t *testing.T) {
		tmpDir := t.TempDir()
		summaryPath := filepath.Join(tmpDir, "summary.md")
		t.Setenv("GITHUB_STEP_SUMMARY", summaryPath)

		p := NewGitHubPresenter()
		err := p.Render(&sampleResult)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		content, err := os.ReadFile(summaryPath)
		if err != nil || len(content) == 0 {
			t.Fatalf("expected non-empty summary file")
		}
	})

	t.Run("Empty Env Var", func(t *testing.T) {
		t.Setenv("GITHUB_STEP_SUMMARY", "")
		p := NewGitHubPresenter()
		_ = p.Render(&sampleResult)
	})
}
