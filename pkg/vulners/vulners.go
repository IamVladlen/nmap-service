package vulners

import (
	"context"
	"strconv"
	"strings"

	"github.com/Ullaakut/nmap/v2"
)

// ScanResult contains
type ScanResult struct {
	Result   *nmap.Run
	Warnings []string
}

// New creates new nmap scanner and runs it with vulners script.
func New(ctx context.Context, targets []string, ports []int32) (*ScanResult, error) {
	var portsStr strings.Builder
	for i, port := range ports {
		if i > 0 {
			portsStr.WriteString(",")
		}
		portsStr.WriteString(strconv.Itoa(int(port)))
	}

	// Create scanner
	sc, err := nmap.NewScanner(
		nmap.WithTargets(targets...),
		nmap.WithPorts(portsStr.String()),
		nmap.WithServiceInfo(),
		nmap.WithContext(ctx),
		nmap.WithScripts("vulners"),
	)
	if err != nil {
		return &ScanResult{}, err
	}

	res, warns, err := sc.Run()

	return &ScanResult{res, warns}, err
}
