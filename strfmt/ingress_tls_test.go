package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseIngressTLS(t *testing.T) {
	r, _ := ParseIngressTLS("secretName:host1,host2,host3")
	require.Equal(t, "secretName", r.SecretName)
	require.Equal(t, "host1", r.Hosts[0])
	require.Equal(t, "host2", r.Hosts[1])
	require.Equal(t, "host3", r.Hosts[2])
}

func TestIngressTLS_String(t *testing.T) {
	ingressTLS := &IngressTLS{
		SecretName: "secretName",
		Hosts:      []string{"host1", "host2", "host3"},
	}

	require.Equal(t, "secretName:host1,host2,host3", ingressTLS.String())
}
