package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseHostAlias(t *testing.T) {
	customHost, err := ParseHostAlias("127.0.0.1 test1.com,test2.com")

	require.Equal(t, nil, err)

	require.Equal(t, "127.0.0.1", customHost.IP)
	require.Equal(t, []string{"test1.com", "test2.com"}, customHost.HostNames)
	require.Equal(t, "127.0.0.1 test1.com,test2.com", customHost.String())
}
