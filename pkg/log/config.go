package log

import (
	"fmt"
	"strings"
)

// Config holds configuration for the logger.
type Config struct {
	Level     string `yaml:"level"`
	Format    string `yaml:"format"`
	AddSource bool   `yaml:"add_source"`
}

// Validate validates the logger configuration.
func (c *Config) Validate() error {
	level := strings.ToLower(c.Level)
	if level != "debug" && level != "info" && level != "warn" && level != "error" {
		return fmt.Errorf("invalid log level: %s", c.Level)
	}

	format := strings.ToLower(c.Format)
	if format != "json" && format != "text" {
		return fmt.Errorf("invalid log format: %s", c.Format)
	}

	return nil
}
