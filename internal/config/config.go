package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Config holds runtime options parsed from CLI flags and environment variables.
type Config struct {
	Seed           string
	Strict         bool
	CustomDeckPath string
}

// Load parses command line arguments and environment variables to construct a Config.
// Priority order: CLI Flags > Environment Variables > Defaults.
func Load(args []string) (*Config, error) {
	flags := flag.NewFlagSet("arcana-gate", flag.ContinueOnError)

	var (
		seed           string
		strict         bool
		customDeckPath string
	)

	flags.StringVar(&seed, "seed", "", "Deterministic seed for divination (defaults to GITHUB_SHA or CI_COMMIT_SHA)")
	flags.BoolVar(&strict, "strict", false, "Enable strict mode where all reversed cards result in a BLOCK")
	flags.StringVar(&customDeckPath, "deck", "", "Path to a custom YAML deck file")

	if err := flags.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse command-line flags: %w", err)
	}

	cfg := &Config{
		Seed:           seed,
		Strict:         strict,
		CustomDeckPath: customDeckPath,
	}

	// Resolve Seed from environment variables if not set via CLI flag
	if cfg.Seed == "" {
		if val := os.Getenv("ARCANA_SEED"); val != "" {
			cfg.Seed = val
		} else if val := os.Getenv("GITHUB_SHA"); val != "" {
			cfg.Seed = val
		} else if val := os.Getenv("CI_COMMIT_SHA"); val != "" {
			cfg.Seed = val
		}
	}

	// Resolve Strict mode from environment variables if false
	if !cfg.Strict {
		if val := os.Getenv("ARCANA_STRICT"); val != "" {
			parsed, err := strconv.ParseBool(val)
			if err == nil {
				cfg.Strict = parsed
			}
		}
	}

	// Resolve CustomDeckPath from environment variables if empty
	if cfg.CustomDeckPath == "" {
		if val := os.Getenv("ARCANA_DECK_PATH"); val != "" {
			cfg.CustomDeckPath = val
		}
	}

	return cfg, nil
}
