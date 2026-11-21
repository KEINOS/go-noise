package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPseudoRrandomSameInput(t *testing.T) {
	first := PseudoRrandom(100, 1.25)
	require.Equal(t, first, PseudoRrandom(100, 1.25))

	PseudoRrandom(100, 1.75)
	PseudoRrandom(101, 1.25)
	require.Equal(t, first, PseudoRrandom(100, 1.25))
}

func TestPseudoRrandomUsesSeedAndCoordinates(t *testing.T) {
	first := PseudoRrandom(100, 1.25)

	require.NotEqual(t, first, PseudoRrandom(101, 1.25))
	require.NotEqual(t, first, PseudoRrandom(100, 1.75))
	require.NotEqual(t, first, PseudoRrandom(100, 1.25, 0))
	require.NotEqual(t, PseudoRrandom(100, 1.25, 2.5), PseudoRrandom(100, 2.5, 1.25))
	require.Equal(t, PseudoRrandom(100, 0), PseudoRrandom(100, math.Copysign(0, -1)))

	for seed := int64(-100); seed <= 100; seed++ {
		value := PseudoRrandom(seed, float64(seed)/4)
		require.Greater(t, value, -1.0)
		require.LessOrEqual(t, value, 1.0)
	}
}
