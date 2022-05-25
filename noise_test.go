package noise_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/KEINOS/go-noise"
	"github.com/stretchr/testify/require"
)

func TestNew_unknown_noise_type(t *testing.T) {
	const seed = 100

	g, err := noise.New(1000, seed)

	require.Error(t, err, "unknown noise type should return an error")
	require.Nil(t, g, "it should be nil on error")
}

func TestNew_is_in_range_perlin(t *testing.T) {
	for i := 0; i < 100; i++ {
		seed := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()

		g, err := noise.New(noise.Perlin, seed)

		require.NoError(t, err, "it should not return an error")
		require.NotNil(t, g, "it should not be nil")

		v := g.Eval64(rand.ExpFloat64())

		require.GreaterOrEqual(t, v, float64(-1), "it should be in the range of -1 to 1")
		require.LessOrEqual(t, v, float64(1), "it should be in the range of -1 to 1")
	}
}

func TestNew_is_in_range_opensimplex(t *testing.T) {
	for i := 0; i < 100; i++ {
		seed := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()

		g, err := noise.New(noise.OpenSimplex, seed)

		require.NoError(t, err, "it should not return an error")
		require.NotNil(t, g, "it should not be nil")

		v := g.Eval64(rand.ExpFloat64())

		require.GreaterOrEqual(t, v, float64(-1), "it should be in the range of -1 to 1")
		require.LessOrEqual(t, v, float64(1), "it should be in the range of -1 to 1")
	}
}
