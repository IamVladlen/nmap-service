package usecase

import (
	"context"
	"testing"

	"github.com/IamVladlen/nmap-service/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type checkVulnArgs struct {
	ctx     context.Context
	targets []string
	ports   []int32
}

func TestUseCase_CheckVuln(t *testing.T) {
	uc := New(logger.New("debug"))

	t.Run("Vulnerable host", func(t *testing.T) {
		args := checkVulnArgs{
			context.Background(),
			[]string{"scanme.nmap.org"},
			[]int32{80},
		}

		got, err := uc.CheckVuln(args.ctx, args.targets, args.ports)
		require.NoError(t, err)
		assert.Equal(t, float32(7.5), got.Results[0].Services[0].Vulns[0].CvssScore) // Find expected vulnerability with 7.5 score
	})

	t.Run("Secure host", func(t *testing.T) {
		args := checkVulnArgs{
			context.Background(),
			[]string{"scanme.nmap.org"},
			[]int32{443},
		}

		got, err := uc.CheckVuln(args.ctx, args.targets, args.ports)
		require.NoError(t, err)
		assert.Equal(t, 0, len(got.Results[0].Services[0].Vulns)) // Check that there is no vulns returned
	})
}
