package config_test

import (
	"testing"

	"github.com/RewithSolo/arcana-gate/internal/config"
)

func TestConfig_CLIFlagsOverride(t *testing.T) {
	t.Setenv("GITHUB_SHA", "env-sha-123")
	t.Setenv("ARCANA_STRICT", "false")

	args := []string{"-seed", "flag-sha-456", "-strict", "-deck", "custom.yaml"}
	cfg, err := config.Load(args)
	if err != nil {
		t.Fatalf("unexpected error loading config: %v", err)
	}

	if cfg.Seed != "flag-sha-456" {
		t.Errorf("expected seed %q from CLI flag, got %q", "flag-sha-456", cfg.Seed)
	}

	if !cfg.Strict {
		t.Errorf("expected strict mode to be true from CLI flag")
	}

	if cfg.CustomDeckPath != "custom.yaml" {
		t.Errorf("expected custom deck path %q, got %q", "custom.yaml", cfg.CustomDeckPath)
	}
}

func TestConfig_EnvFallbackChain(t *testing.T) {
	t.Run("ARCANA_SEED takes precedence over GITHUB_SHA", func(t *testing.T) {
		t.Setenv("ARCANA_SEED", "arcana-seed-value")
		t.Setenv("GITHUB_SHA", "github-sha-value")

		cfg, err := config.Load([]string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.Seed != "arcana-seed-value" {
			t.Errorf("expected seed %q, got %q", "arcana-seed-value", cfg.Seed)
		}
	})

	t.Run("CI_COMMIT_SHA fallback", func(t *testing.T) {
		t.Setenv("ARCANA_SEED", "")
		t.Setenv("GITHUB_SHA", "")
		t.Setenv("CI_COMMIT_SHA", "gitlab-commit-sha")

		cfg, err := config.Load([]string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.Seed != "gitlab-commit-sha" {
			t.Errorf("expected seed %q, got %q", "gitlab-commit-sha", cfg.Seed)
		}
	})

	t.Run("ARCANA_DECK_PATH env fallback", func(t *testing.T) {
		t.Setenv("ARCANA_DECK_PATH", "/path/to/deck.yaml")

		cfg, err := config.Load([]string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.CustomDeckPath != "/path/to/deck.yaml" {
			t.Errorf("expected custom deck path %q, got %q", "/path/to/deck.yaml", cfg.CustomDeckPath)
		}
	})

	t.Run("Invalid ARCANA_STRICT bool handled gracefully", func(t *testing.T) {
		t.Setenv("ARCANA_STRICT", "invalid-boolean")

		cfg, err := config.Load([]string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.Strict {
			t.Errorf("expected strict to remain false on invalid bool string")
		}
	})
}

func TestConfig_InvalidFlagError(t *testing.T) {
	args := []string{"-unknown-flag"}
	_, err := config.Load(args)
	if err == nil {
		t.Fatal("expected error when parsing unknown flag, got nil")
	}
}
