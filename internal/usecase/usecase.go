package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IamVladlen/nmap-service/pkg/logger"
	"github.com/IamVladlen/nmap-service/pkg/vulners"
	grpcnmap "github.com/IamVladlen/nmap-service/proto/nmap"
	"github.com/Ullaakut/nmap/v2"
)

type UseCase struct {
	log *logger.Log
}

// CheckVuln scans hosts and detects their vulnerabilities.
func (uc *UseCase) CheckVuln(ctx context.Context, targets []string, ports []int32) (*grpcnmap.CheckVulnResponse, error) {
	result, err := vulners.New(ctx, targets, ports)
	for _, warn := range result.Warnings {
		uc.log.Warn().Str("Warning", warn).Msg("Received warning while scanning network")
	}
	if err != nil {
		return &grpcnmap.CheckVulnResponse{}, fmt.Errorf("usecase - CheckVuln: %w", err)
	}

	response := parseResult(result)

	return response, nil
}

// parseResult binds scan results to response struct.
func parseResult(res *vulners.ScanResult) *grpcnmap.CheckVulnResponse {
	var response grpcnmap.CheckVulnResponse

	// Bind to CheckVulnResponse.TargetResult
	for _, h := range res.Result.Hosts {
		host := grpcnmap.TargetResult{
			Target: h.Addresses[0].String(),
		}

		// Bind to CheckVulnResponse.TargetResult.Service
		for _, p := range h.Ports {
			port := grpcnmap.Service{
				Name:    p.Service.Name,
				Version: p.Service.Version,
				TcpPort: int32(p.ID),
			}

			// Bind to CheckVulnResponse.TargetResult.Service.Vulnerability
			for _, script := range p.Scripts {
				if script.ID == "vulners" {
					getVulns(&port, script)
				}
			}

			host.Services = append(host.Services, &port)
		}

		response.Results = append(response.Results, &host)
	}

	return &response
}

func getVulns(port *grpcnmap.Service, script nmap.Script) {
	for _, v := range script.Tables[0].Tables {
		cvss, err := strconv.ParseFloat(findElement(v.Elements, "cvss"), 32)
		if err != nil {
			continue
		}

		vuln := grpcnmap.Vulnerability{
			Identifier: findElement(v.Elements, "id"),
			CvssScore:  float32(cvss),
		}

		port.Vulns = append(port.Vulns, &vuln)
	}
}

func findElement(table []nmap.Element, key string) string {
	for _, el := range table {
		if el.Key == key {
			return el.Value
		}
	}

	return ""
}

func New(log *logger.Log) *UseCase {
	return &UseCase{
		log: log,
	}
}
