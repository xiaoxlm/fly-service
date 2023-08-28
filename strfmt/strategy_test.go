package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrategy(t *testing.T) {
	strategy, err := ParseStrategy("RollingUpdate:maxUnavailable=25%,maxSurge=25%")
	require.Equal(t, nil, err)

	require.Equal(t, "RollingUpdate", strategy.Type)
	require.Equal(t, map[string]string{
		"maxUnavailable": "25%",
		"maxSurge":       "25%",
	}, strategy.Flags)
	require.Equal(t, "RollingUpdate:maxSurge=25%,maxUnavailable=25%", strategy.String())
}
