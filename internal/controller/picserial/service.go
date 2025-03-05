package picserial

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/tbe-team/raybot/internal/controller/picserial/handler"
	"github.com/tbe-team/raybot/internal/controller/picserial/serial"
	"github.com/tbe-team/raybot/internal/service"
)

type Config struct {
	Serial serial.Config `yaml:"serial"`
}

// RegisterFlags registers flags for the PIC configuration.
func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	cfg.Serial.RegisterFlagsWithPrefix("pic-", f)
}

// Validate validates the PIC configuration.
func (cfg *Config) Validate() error {
	return cfg.Serial.Validate()
}

type Handlers struct {
	SyncStateHandler *handler.SyncStateHandler
}

//nolint:revive
type PICSerialService struct {
	cfg Config

	serialClient serial.Client

	handlers Handlers
	log      *slog.Logger
}

type CleanupFunc func(context.Context) error

func NewPICSerialService(cfg Config, service service.Service) (*PICSerialService, error) {
	serialClient, err := serial.NewClient(cfg.Serial)
	if err != nil {
		return nil, fmt.Errorf("failed to create serial client: %w", err)
	}

	handlers := Handlers{
		SyncStateHandler: handler.NewSyncStateHandler(service.RobotService()),
	}

	return &PICSerialService{
		cfg:          cfg,
		serialClient: serialClient,
		handlers:     handlers,
		log:          slog.With(slog.String("service", "PICSerialService")),
	}, nil
}

// Run runs the PIC serial service.
func (s *PICSerialService) Run(ctx context.Context) (CleanupFunc, error) {
	s.log.Debug("PIC serial service is running")

	go s.readLoop(ctx)

	cleanup := func(_ context.Context) error {
		s.log.Debug("PIC serial service is shutting down")
		if err := s.serialClient.Stop(); err != nil {
			s.log.Error("failed to stop serial client", "error", err)
			return err
		}
		s.log.Debug("PIC serial service shut down complete")
		return nil
	}

	return cleanup, nil
}

func (s *PICSerialService) readLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-s.serialClient.Read():
			if !ok {
				s.log.Error("serial client read channel closed")
				return
			}
			s.routeMessage(ctx, msg)
		}
	}
}

func (s *PICSerialService) routeMessage(ctx context.Context, msg []byte) {
	var temp struct {
		Type messageType `json:"type"`
	}
	if err := json.Unmarshal(msg, &temp); err != nil {
		s.log.Error("failed to unmarshal message", slog.Any("error", err), slog.Any("message", msg))
		return
	}

	//nolint:gocritic
	switch temp.Type {
	case messageTypeSyncState:
		var syncStateMsg handler.SyncStateMessage
		if err := json.Unmarshal(msg, &syncStateMsg); err != nil {
			s.log.Error("failed to unmarshal sync state message", slog.Any("error", err), slog.Any("message", msg))
			return
		}
		s.handlers.SyncStateHandler.Handle(ctx, syncStateMsg)
	}
}

// messageType is the type of message received from the PIC
type messageType uint8

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *messageType) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseUint(string(data), 10, 8)
	if err != nil {
		return fmt.Errorf("parse uint8: %w", err)
	}

	switch n {
	case 0:
		*m = messageTypeSyncState
	default:
		return fmt.Errorf("invalid message type: %s", string(data))
	}
	return nil
}

const (
	messageTypeSyncState messageType = iota
)
