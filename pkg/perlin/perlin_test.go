package perlin_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-noise/pkg/perlin"
	goperlin "github.com/aquilax/go-perlin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_original_perlin(t *testing.T) {
	const seed = 100

	p := goperlin.NewPerlin(perlin.Alpha, perlin.Beta, perlin.Iteration, seed)

	for _, test := range []struct {
		input  float64
		expect string
	}{
		{0, "0;0.0000"},
		{1, "1;-0.0086"},
		{2, "2;-0.0017"},
	} {
		x := test.input
		expect := test.expect
		actual := fmt.Sprintf("%0.0f;%0.4f", x, p.Noise1D(x/10))

		assert.Equal(t, expect, actual)
	}
}

func TestNew_instantiate(t *testing.T) {
	const seed = 100

	p := perlin.New(seed)

	require.Equal(t, 2., p.Smoothness, "the Smoothness should be float 2 by default")
	require.Equal(t, 2., p.Scale, "the Scale should be float 2 by default")
	require.Equal(t, int64(seed), p.Seed, "the Seed should be int64(seed) by default")
	require.Equal(t, perlin.Iteration, p.Iteration, "the Iteration should be perlin.Iteration by default")
}
