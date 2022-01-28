package bss_test

import (
	"fmt"
	"testing"

	"github.com/ethereum-optimism/optimistic-specs/opnode/bss"
	"github.com/stretchr/testify/require"
)

var validateConfigTests = []struct {
	name   string
	cfg    bss.Config
	expErr error
}{
	{
		name: "bad log level",
		cfg: bss.Config{
			LogLevel: "unknown",
		},
		expErr: fmt.Errorf("unknown level: unknown"),
	},
	// Valid configs
	{
		name: "valid log level",
		cfg: bss.Config{
			LogLevel: "info",
		},
		expErr: nil,
	},
}

// TestValidateConfig asserts the behavior of ValidateConfig by testing expected
// error and success configurations.
func TestValidateConfig(t *testing.T) {
	for _, test := range validateConfigTests {
		t.Run(test.name, func(t *testing.T) {
			err := bss.ValidateConfig(&test.cfg)
			require.Equal(t, err, test.expErr)
		})
	}
}
