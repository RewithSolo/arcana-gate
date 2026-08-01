package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/RewithSolo/arcana-gate/internal/config"
	"github.com/RewithSolo/arcana-gate/internal/domain"
	"github.com/RewithSolo/arcana-gate/internal/engine"
	"github.com/RewithSolo/arcana-gate/internal/presenter"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\n❌ %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// 1. Parse configuration from flags and environment variables
	cfg, err := config.Load(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Fallback seed if none was provided by flags, env, or CI context
	if cfg.Seed == "" {
		cfg.Seed = fmt.Sprintf("runtime-seed-%d", time.Now().UnixNano())
	}

	// 2. Instantiate the divination engine
	// Default deck will be used if customDeck is empty
	var customDeck []domain.Card
	eng := engine.New(customDeck, engine.Config{
		Strict: cfg.Strict,
	})

	// 3. Execute divination logic with timeout context safety
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := eng.Divine(ctx, cfg.Seed)
	if err != nil {
		return fmt.Errorf("divination failed: %w", err)
	}

	// 4. Render output to Terminal (stdout)
	termPresenter := presenter.NewTerminalPresenter(os.Stdout)
	if err := termPresenter.Render(result); err != nil {
		return fmt.Errorf("failed to render terminal output: %w", err)
	}

	// 5. Render output to GitHub Actions Summary (if running in GitHub environment)
	ghPresenter := presenter.NewGitHubPresenter()
	if err := ghPresenter.Render(result); err != nil {
		// Non-fatal error: log warning to stderr but do not break the pipeline execution
		fmt.Fprintf(os.Stderr, "Warning: failed to write GitHub step summary: %v\n", err)
	}

	// 6. Enforce gate status exit policy
	if result.Status == domain.StatusBlock {
		return fmt.Errorf("arcana gate blocked the pipeline: %s", result.Reason)
	}

	return nil
}
