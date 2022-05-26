/*
Package custom provides a noise generator using the user-defined function.
Which implements github.com/KEINOS/go-noise/noise interface.

You need to define the function before calling `Eval32` or `Eval64`.
*/
package custom

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// Generator holds parameters of user-defined noise generator.
type Generator struct {
	// func32 is the user-defined function for float32.
	func32 func(seed int64, dim ...float32) float32
	// func64 is the user-defined function for float64.
	func64 func(seed int64, dim ...float64) float64
	// Seed holds the seed value for the noise.
	Seed int64
}

// New returns a seeded Perlin noise instance.
func New(seed int64) *Generator {
	// Rand that uses random values from src to generate other random values.
	return &Generator{
		Seed: seed,
	}
}

// Eval32 returns a float32 noise value using the user-defined function for the given coordinates.
func (n *Generator) Eval32(dim ...float32) float32 {
	return n.func32(n.Seed, dim...)
}

// Eval64 returns a float64 noise value using the user-defined function for the given coordinates.
func (n *Generator) Eval64(dim ...float64) float64 {
	return n.func64(n.Seed, dim...)
}

// SetEval32 sets the user-defined noise generator function for float32.
func (n *Generator) SetEval32(f func(seed int64, dim ...float32) float32) error {
	// Check the function.
	rand.Seed(time.Now().UnixNano())

	//nolint:gosec // Use of weak random number generation is intended here.
	dummyInput := rand.Float32()

	if v := f(n.Seed, dummyInput, dummyInput); v <= -1 || 1 < v {
		return errors.New("the function must return a value between -1 and 1")
	}

	n.func32 = f

	return nil
}

// SetEval64 sets the user-defined noise generator function for float64.
func (n *Generator) SetEval64(f func(seed int64, dim ...float64) float64) error {
	// Check the function.
	rand.Seed(time.Now().UnixNano())

	//nolint:gosec // Use of weak random number generation is intended here.
	dummyInput := rand.Float64()

	if v := f(n.Seed, dummyInput, dummyInput); v <= -1 || 1 < v {
		return errors.New("the function must return a value between -1 and 1")
	}

	n.func64 = f

	return nil
}
