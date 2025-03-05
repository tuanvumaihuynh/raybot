package grpc

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"runtime/debug"

	"buf.build/gen/go/tbe-team/raybot-api/grpc/go/raybot/v1/raybotv1grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tbe-team/raybot/internal/controller/grpc/handler"
	"github.com/tbe-team/raybot/internal/service"
)

type Config struct {
	Port int `yaml:"port"`
}

func (c *Config) RegisterFlags(f *flag.FlagSet) {
	f.IntVar(&c.Port, "grpc-port", 50051, "gRPC port")
}

func (c *Config) Validate() error {
	if c.Port < 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	return nil
}

type CleanupFunc func(context.Context) error

//nolint:revive
type GRPCService struct {
	cfg Config

	service service.Service
	log     *slog.Logger
}

func NewGRPCService(cfg Config, service service.Service) (*GRPCService, error) {
	return &GRPCService{
		cfg:     cfg,
		service: service,
		log:     slog.With(slog.String("service", "GRPCService")),
	}, nil
}

func (s GRPCService) Run() (CleanupFunc, error) {
	interceptorLogger := func(l *slog.Logger) logging.Logger {
		return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l.Log(ctx, slog.Level(lvl), msg, fields...)
		})
	}
	grpcPanicRecoveryHandler := func(p any) (err error) {
		s.log.Error("recovered from panic", "panic", p, "stack", debug.Stack())
		return status.Errorf(codes.Internal, "%s", p)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(s.log)),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(s.log)),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)

	s.registerHandlers(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return nil, err
	}

	go func() {
		s.log.Debug(fmt.Sprintf("GRPC server is listening on port %d", s.cfg.Port))
		if err := server.Serve(lis); err != nil {
			s.log.Error("failed to serve GRPC", "error", err)
		}
	}()

	cleanup := func(_ context.Context) error {
		s.log.Debug("GRPC server is shutting down")
		server.Stop()
		s.log.Debug("GRPC server shut down complete")
		return nil
	}

	return cleanup, nil
}

func (s GRPCService) registerHandlers(server *grpc.Server) {
	robotStateHandler := handler.NewRobotStateHandler(s.service.RobotService())
	raybotv1grpc.RegisterRobotStateServiceServer(server, robotStateHandler)
}
