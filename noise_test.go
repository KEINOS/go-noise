package noise_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/KEINOS/go-noise"
	"github.com/stretchr/testify/require"
)

func TestNew_unknown_noise_type(t *testing.T) {
	t.Parallel()

	const seed = 100

	for _, algorithm := range []noise.Algo{noise.Unknown, 1000} {
		g, err := noise.New(algorithm, seed)

		require.Error(t, err, "unknown noise type %d should return an error", algorithm)
		require.Nil(t, g, "unknown noise type %d should return no generator", algorithm)
	}
}

func TestNew_is_in_range_perlin(t *testing.T) {
	t.Parallel()

	for range 100 {
		//nolint:gosec // Use of weak random number generator is OK here
		seed := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()

		g, err := noise.New(noise.Perlin, seed)

		require.NoError(t, err, "it should not return an error")
		require.NotNil(t, g, "it should not be nil")

		v := g.Eval64(1.0)

		require.GreaterOrEqual(t, v, float64(-1), "it should be in the range of -1 to 1")
		require.LessOrEqual(t, v, float64(1), "it should be in the range of -1 to 1")
	}
}

func TestNew_is_in_range_opensimplex(t *testing.T) {
	t.Parallel()

	for range 100 {
		//nolint:gosec // Use of weak random number generator is OK here
		seed := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()

		g, err := noise.New(noise.OpenSimplex, seed)

		require.NoError(t, err, "it should not return an error")
		require.NotNil(t, g, "it should not be nil")

		v := g.Eval64(1.0)

		require.GreaterOrEqual(t, v, float64(-1), "it should be in the range of -1 to 1")
		require.LessOrEqual(t, v, float64(1), "it should be in the range of -1 to 1")
	}
}
