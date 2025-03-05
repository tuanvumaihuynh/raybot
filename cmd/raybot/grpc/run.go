package grpc

import (
	"context"
	"fmt"

	"github.com/tbe-team/raybot/internal/application"
	"github.com/tbe-team/raybot/internal/controller/grpc"
)

func Start(app *application.Application) error {
	grpcService, err := grpc.NewGRPCService(app.Cfg.GRPC, app.Service)
	if err != nil {
		return fmt.Errorf("failed to create GRPC service: %w", err)
	}

	cleanup, err := grpcService.Run()
	if err != nil {
		return fmt.Errorf("failed to run GRPC service: %w", err)
	}

	app.CleanupManager.Add(func(ctx context.Context) error {
		return cleanup(ctx)
	})

	return nil
}
