package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseAction(t *testing.T) {
	t.Run("#tcp default", func(t *testing.T) {
		action, err := ParseAction("tcp://:22")
		require.Equal(t, nil, err)
		require.Equal(t, "22", action.TCPSocket.Port.String())
		require.Equal(t, "", action.TCPSocket.Host)

	})
	t.Run("#http default", func(t *testing.T) {
		action, err := ParseAction("http://:8080")
		require.Equal(t, nil, err)

		require.Equal(t, "HTTP", string(action.HTTPGet.Scheme))
		require.Equal(t, "8080", action.HTTPGet.Port.String())
	})
	t.Run("#http url", func(t *testing.T) {
		action, err := ParseAction("http://127.0.0.1:8080/health")
		require.Equal(t, nil, err)

		require.Equal(t, "HTTP", string(action.HTTPGet.Scheme))
		require.Equal(t, "127.0.0.1", action.HTTPGet.Host)
		require.Equal(t, "8080", action.HTTPGet.Port.String())
		require.Equal(t, "/health", action.HTTPGet.Path)
	})
	t.Run("#exec", func(t *testing.T) {
		action, err := ParseAction("touch /tmp/healthy; sleep 30; rm -f /tmp/healthy; sleep 600")
		require.Equal(t, nil, err)

		require.Equal(t, []string{"sh", "-c", "touch /tmp/healthy; sleep 30; rm -f /tmp/healthy; sleep 600"}, action.Exec.Command)
	})
}
