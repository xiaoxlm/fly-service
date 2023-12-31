package strfmt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequestAndLimit(t *testing.T) {
	t.Run("#parse & string", func(t *testing.T) {
		r, err := ParseRequestLimit("1/500")
		require.Equal(t, nil, err)

		require.Equal(t, 1, r.Request)
		require.Equal(t, 500, r.Limit)
		require.Equal(t, r.String(), "1/500")

	})

	t.Run("#parse & string with unit", func(t *testing.T) {
		r, err := ParseRequestLimit("10/500e6")
		require.Equal(t, nil, err)

		require.Equal(t, 10, r.Request)
		require.Equal(t, 500, r.Limit)
		require.Equal(t, "e6", r.Unit)
		require.Equal(t, "10/500e6", r.String())
	})

	t.Run("#parse & string simple", func(t *testing.T) {
		r, err := ParseRequestLimit("10")
		require.Equal(t, nil, err)

		require.Equal(t, 10, r.Request)
		require.Equal(t, 0, r.Limit)
		require.Equal(t, "10", r.String())
	})
}
