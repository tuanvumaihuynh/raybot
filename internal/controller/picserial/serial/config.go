package serial

import (
	"fmt"
	"time"
)

// Config is the configuration for the PIC serial port.
type Config struct {
	// Port is the serial port path (e.g., "/dev/ttyUSB0" or "COM3")
	Port string `yaml:"port"`

	// BaudRate is the communication speed
	BaudRate int `yaml:"baud_rate"`

	// DataBits is the number of data bits (usually 8)
	DataBits int `yaml:"data_bits"`

	// StopBits is the number of stop bits (usually 1)
	StopBits float64 `yaml:"stop_bits"`

	// Parity mode (none, odd, even)
	Parity string `yaml:"parity"`

	// ReadTimeout is the timeout for read operations
	ReadTimeout time.Duration `yaml:"read_timeout"`
}

// Validate verifies the configuration for the PIC serial port.
func (cfg *Config) Validate() error {
	if cfg.BaudRate < 1200 || cfg.BaudRate > 115200 {
		return fmt.Errorf("invalid baud rate: %d", cfg.BaudRate)
	}

	if cfg.DataBits != 5 && cfg.DataBits != 6 && cfg.DataBits != 7 && cfg.DataBits != 8 {
		return fmt.Errorf("invalid data bits: %d", cfg.DataBits)
	}

	if cfg.StopBits != 1 && cfg.StopBits != 1.5 && cfg.StopBits != 2 {
		return fmt.Errorf("invalid stop bits: %f", cfg.StopBits)
	}

	if cfg.Parity != "none" && cfg.Parity != "odd" && cfg.Parity != "even" {
		return fmt.Errorf("invalid parity: %s", cfg.Parity)
	}

	return nil
}
