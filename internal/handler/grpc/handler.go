package grpchandler

import (
	"context"

	"github.com/IamVladlen/nmap-service/internal/usecase"
	"github.com/IamVladlen/nmap-service/pkg/grpcserver"
	"github.com/IamVladlen/nmap-service/pkg/logger"
	grpcnmap "github.com/IamVladlen/nmap-service/proto/nmap"
)

type handler struct {
	grpcnmap.UnimplementedNetVulnServiceServer

	uc  *usecase.UseCase
	log *logger.Log
}

// CheckVuln scans hosts and detects their vulnerabilities.
func (h *handler) CheckVuln(ctx context.Context, req *grpcnmap.CheckVulnRequest) (*grpcnmap.CheckVulnResponse, error) {
	h.log.Info().Str("Endpoint", "CheckVuln").Msg("Host scan started")

	resp, err := h.uc.CheckVuln(ctx, req.GetTargets(), req.GetTcpPort())
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to check vulnerabilities")

		return &grpcnmap.CheckVulnResponse{}, err
	}

	h.log.Info().Str("Endpoint", "CheckVuln").Msg("Host scan ended successfully")

	return resp, nil
}

func New(srv *grpcserver.Server, uc *usecase.UseCase, log *logger.Log) {
	h := &handler{
		uc:  uc,
		log: log,
	}

	grpcnmap.RegisterNetVulnServiceServer(srv, h)
}
