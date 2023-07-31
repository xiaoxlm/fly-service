package strfmt

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestParseVolumeMount(t *testing.T) {
	t.Run("parse & string simple", func(t *testing.T) {
		r, err := ParseVolumeMount("data:/html")
		require.Equal(t, nil, err)

		require.Equal(t, "data", r.Name)
		require.Equal(t, "/html", r.MountPath)
		require.Equal(t, false, r.ReadOnly)
		require.Equal(t, "", r.SubPath)
		require.Equal(t, "data:/html", r.String())
	})

	t.Run("parse & string", func(t *testing.T) {
		r, err := ParseVolumeMount("data/html:/html:ro")
		require.Equal(t, nil, err)

		require.Equal(t, "data", r.Name)
		require.Equal(t, "/html", r.MountPath)
		require.Equal(t, true, r.ReadOnly)
		require.Equal(t, "html", r.SubPath)
		require.Equal(t, "data/html:/html:ro", r.String())
	})

	t.Run("VolumeMount yaml marshal & unmarshal", func(t *testing.T) {
		data, err := yaml.Marshal(struct {
			Mount *VolumeMount `yaml:"volumeMount"`
		}{
			Mount: &VolumeMount{
				MountPath: "/html",
				Name:      "data",
				ReadOnly:  true,
				SubPath:   "html",
			},
		})

		require.Equal(t, nil, err)
		require.Equal(t, "volumeMount: data/html:/html:ro\n", string(data))

		v := struct {
			Mount VolumeMount `yaml:"volumeMount"`
		}{}

		err = yaml.Unmarshal(data, &v)

		require.Equal(t, nil, err)
		require.Equal(t, "data/html:/html:ro", v.Mount.String())
	})
}
