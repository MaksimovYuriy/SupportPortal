package config

import (
	"testing"
	"time"
)

func TestLoadEngineDefaults(t *testing.T) {
	t.Setenv("ENGINE_ENABLED", "")
	t.Setenv("ENGINE_INTERVAL", "")
	t.Setenv("ENGINE_BATCH_LIMIT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if !cfg.Engine.Enabled {
		t.Fatalf("expected engine enabled by default")
	}
	if cfg.Engine.Interval != 5*time.Second {
		t.Fatalf("expected default interval 5s, got %s", cfg.Engine.Interval)
	}
	if cfg.Engine.BatchLimit != 100 {
		t.Fatalf("expected default batch limit 100, got %d", cfg.Engine.BatchLimit)
	}
}

func TestLoadEngineOverrides(t *testing.T) {
	t.Setenv("ENGINE_ENABLED", "false")
	t.Setenv("ENGINE_INTERVAL", "250ms")
	t.Setenv("ENGINE_BATCH_LIMIT", "7")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.Engine.Enabled {
		t.Fatalf("expected engine disabled")
	}
	if cfg.Engine.Interval != 250*time.Millisecond {
		t.Fatalf("expected interval 250ms, got %s", cfg.Engine.Interval)
	}
	if cfg.Engine.BatchLimit != 7 {
		t.Fatalf("expected batch limit 7, got %d", cfg.Engine.BatchLimit)
	}
}

func TestLoadEngineInvalidValuesUseDefaults(t *testing.T) {
	t.Setenv("ENGINE_ENABLED", "maybe")
	t.Setenv("ENGINE_INTERVAL", "soon")
	t.Setenv("ENGINE_BATCH_LIMIT", "many")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if !cfg.Engine.Enabled {
		t.Fatalf("expected invalid enabled value to fall back to true")
	}
	if cfg.Engine.Interval != 5*time.Second {
		t.Fatalf("expected invalid interval to fall back to 5s, got %s", cfg.Engine.Interval)
	}
	if cfg.Engine.BatchLimit != 100 {
		t.Fatalf("expected invalid batch limit to fall back to 100, got %d", cfg.Engine.BatchLimit)
	}
}
