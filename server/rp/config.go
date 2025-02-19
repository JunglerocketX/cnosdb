package rp

import (
	"errors"
	"github.com/cnosdb/cnosdb/vend/common/monitor/diagnostics"
	"time"

	"github.com/cnosdb/cnosdb/vend/common/pkg/toml"
)

// Config represents the configuration for the rp service.
type Config struct {
	Enabled       bool          `toml:"enabled"`
	CheckInterval toml.Duration `toml:"check-interval"`
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() Config {
	return Config{Enabled: true, CheckInterval: toml.Duration(30 * time.Minute)}
}

// Validate returns an error if the Config is invalid.
func (c Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	// TODO: Should we enforce a minimum interval?
	// Polling every nanosecond, for instance, will greatly impact performance.
	if c.CheckInterval <= 0 {
		return errors.New("check-interval must be positive")
	}

	return nil
}

// Diagnostics returns a diagnostics representation of a subset of the Config.
func (c Config) Diagnostics() (*diagnostics.Diagnostics, error) {
	if !c.Enabled {
		return diagnostics.RowFromMap(map[string]interface{}{
			"enabled": false,
		}), nil
	}

	return diagnostics.RowFromMap(map[string]interface{}{
		"enabled":        true,
		"check-interval": c.CheckInterval,
	}), nil
}
