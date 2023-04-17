package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/IamVladlen/nmap-service/config"
	"github.com/IamVladlen/nmap-service/pkg/grpcserver"
	"github.com/IamVladlen/nmap-service/pkg/logger"

	grpchandler "github.com/IamVladlen/nmap-service/internal/handler/grpc"
	"github.com/IamVladlen/nmap-service/internal/usecase"
)

// Run injects dependencies and starts gRPC server.
func Run(cfg *config.Config, log *logger.Log) {
	uc := usecase.New(log)

	grpcSrv, err := grpcserver.New(cfg.GRPC.PORT, log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start gRPC server")
	}
	grpchandler.New(grpcSrv, uc, log)

	log.Info().Msgf("Service successfully started on port %s", cfg.GRPC.PORT)

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-sigCh:
		log.Info().Msgf("Server is shutting down: %s signal", s.String())
	case err := <-grpcSrv.Notify():
		log.Error().Err(err).Msg("Server is shutting down due to error occurrence")
	}

	grpcSrv.Stop()
}
