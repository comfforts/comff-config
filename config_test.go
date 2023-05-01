package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	config "github.com/comfforts/comff-config"
)

const (
	ADDR = "localhost:55055"
)

func TestConfig(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
	){
		"test server config, succeeds":          testServerConfig,
		"test client config, succeeds":          testClientConfig,
		"test custom config error, succeeds":    testCustomConfigError,
		"test undefined target error, succeeds": testUndefinedTargetError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testServerConfig(t *testing.T) {
	t.Helper()

	opts := config.ConfigOpts{
		Addr:   ADDR,
		Target: config.SERVER,
	}

	cfg, err := config.SetupTLSConfig(&opts)
	require.NoError(t, err)
	require.Equal(t, cfg.ServerName, ADDR)
	// fmt.Print("TLS Config: ", cfg)
}

func testClientConfig(t *testing.T) {
	t.Helper()

	opts := config.ConfigOpts{
		Target: config.CLIENT,
	}

	cfg, err := config.SetupTLSConfig(&opts)
	require.NoError(t, err)
	require.Equal(t, cfg.ServerName, "")
}

func testCustomConfigError(t *testing.T) {
	t.Helper()

	opts := config.ConfigOpts{
		Addr:   ADDR,
		Target: config.CUSTOM,
	}

	_, err := config.SetupTLSConfig(&opts)
	require.Error(t, err)
	require.Equal(t, err, config.ErrMissingRequired)
}

func testUndefinedTargetError(t *testing.T) {
	t.Helper()

	opts := config.ConfigOpts{
		Addr: ADDR,
	}

	_, err := config.SetupTLSConfig(&opts)
	require.Error(t, err)
	require.Equal(t, err, config.ErrUndefinedTarget)
}
