package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

func TestPortForward(t *testing.T) {
	t.Run("parse & string", func(t *testing.T) {
		portForward, err := ParsePortForward("80:8080/TCP")
		require.Equal(t, nil, err)

		require.Equal(t, uint16(80), portForward.Port)
		require.Equal(t, uint16(8080), portForward.TargetPort)
		require.Equal(t, "TCP", portForward.Protocol)
		require.Equal(t, "80:8080/TCP", portForward.String())
	})

	t.Run("parse & string without target port ", func(t *testing.T) {
		portForward, err := ParsePortForward("80/TCP")
		require.Equal(t, nil, err)

		require.Equal(t, uint16(80), portForward.Port)
		require.Equal(t, uint16(80), portForward.TargetPort)
		require.Equal(t, "TCP", portForward.Protocol)
		require.Equal(t, "80:80/TCP", portForward.String())
	})

	t.Run("parse & string without protocol", func(t *testing.T) {
		portForward, err := ParsePortForward("80:8080")
		require.Equal(t, nil, err)

		require.Equal(t, uint16(80), portForward.Port)
		require.Equal(t, uint16(8080), portForward.TargetPort)
		require.Equal(t, "80:8080/TCP", portForward.String())
	})

	t.Run("yaml marshal & unmarshal", func(t *testing.T) {
		data, err := yaml.Marshal(struct {
			Port *PortForward `yaml:"port"`
		}{
			Port: &PortForward{
				Port:       80,
				TargetPort: 8080,
				Protocol:   "TCP",
			},
		})
		require.Equal(t, nil, err)
		require.Equal(t, "port: 80:8080/TCP\n", string(data))

		v := struct {
			Port PortForward `yaml:"port"`
		}{}

		err = yaml.Unmarshal(data, &v)
		require.Equal(t, nil, err)
		NewWithT(t).Expect(v.Port.String()).To(Equal("80:8080/TCP"))
	})
}
