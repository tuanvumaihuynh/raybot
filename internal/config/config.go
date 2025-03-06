package config

import (
	"fmt"
	"time"

	"github.com/tbe-team/raybot/internal/controller/grpc"
	"github.com/tbe-team/raybot/internal/controller/picserial"
	"github.com/tbe-team/raybot/internal/controller/picserial/serial"
	"github.com/tbe-team/raybot/pkg/log"
)

type Config struct {
	PIC  picserial.Config `yaml:"pic"`
	GRPC grpc.Config      `yaml:"grpc"`
	Log  log.Config       `yaml:"log"`

	ConfigPath string `yaml:"-"`
}

// Validate validates the application configuration.
func (cfg *Config) Validate() error {
	if err := cfg.PIC.Validate(); err != nil {
		return fmt.Errorf("validate PIC: %w", err)
	}

	if err := cfg.GRPC.Validate(); err != nil {
		return fmt.Errorf("validate GRPC: %w", err)
	}

	if err := cfg.Log.Validate(); err != nil {
		return fmt.Errorf("validate log: %w", err)
	}

	return nil
}

// DefaultConfig is the default configuration for the application.
var DefaultConfig = Config{
	PIC: picserial.Config{
		Serial: serial.Config{
			Port:        "/dev/ttyUSB0",
			BaudRate:    9600,
			DataBits:    8,
			Parity:      "none",
			StopBits:    1,
			ReadTimeout: 1 * time.Second,
		},
	},
	GRPC: grpc.Config{
		Port: 50051,
	},
	Log: log.Config{
		Level:     "info",
		Format:    "json",
		AddSource: false,
	},
}

func init() {
	// Ensure the default config is valid
	if err := DefaultConfig.Validate(); err != nil {
		panic(err)
	}
}
