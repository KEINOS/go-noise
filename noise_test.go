package noise_test

import (
	"testing"

	"github.com/KEINOS/go-noise"
	"github.com/stretchr/testify/require"
)

func TestNew_unknown_noise_type(t *testing.T) {
	const seed = 100

	g, err := noise.New(1000, seed)

	require.Error(t, err, "unknown noise type should return an error")
	require.Nil(t, g, "it should be nil on error")
}
