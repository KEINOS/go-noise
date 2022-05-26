/*
Package perlin is a wrapper of github.com/aquilax/go-perlin.

Which implements github.com/KEINOS/go-noise/noise interface.
*/
package perlin

import (
	"github.com/pkg/errors"

	goperlin "github.com/aquilax/go-perlin"
)

const (
	// Alpha is the default smoothness of Perlin noise.
	Alpha = 2.
	// Beta is the default scale of Perlin noise.
	Beta = 2.
	// Iteration is the default number of iterations of Perlin noise.
	Iteration = int32(3)
)

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a seeded Perlin noise instance.
func New(seed int64) *Noise {
	return &Noise{
		Smoothness: Alpha,
		Scale:      Beta,
		Seed:       seed,
		Iteration:  Iteration,
	}
}

// ----------------------------------------------------------------------------
//  Type: Noise
// ----------------------------------------------------------------------------

// Noise holds parameters of Perlin noise.
type Noise struct {
	// Smoothness is the weight when the sum is formed. Which is the alpha value
	// in Ken Perlin's article. Default is 2. The smaller the number, the more
	// noise is generated.
	Smoothness float64
	// Scale is the harmonic scaling/spacing, typically 2. Which is the beta value
	// in Ken Perlin's article. Default is 2.
	Scale float64
	// Iteration is the number of times the noise is generated. Which is the number
	// of iterations in Ken Perlin's article. Default is 3.
	Iteration int32
	// Seed holds the seed value for the noise.
	Seed int64
}

// ----------------------------------------------------------------------------
//  Methods (Public)
// ----------------------------------------------------------------------------

// Eval32 returns a float32 Perlin noise value for the given coordinates.
// It is a conversion of float64 to float32 to support Eval32 interface.
func (n *Noise) Eval32(dim ...float32) float32 {
	switch len(dim) {
	case 1:
		return float32(n.eval1D(float64(dim[0])))
	case 2:
		return float32(n.eval2D(float64(dim[0]), float64(dim[1])))
	case 3:
		return float32(n.eval3D(float64(dim[0]), float64(dim[1]), float64(dim[2])))
	}

	return 0
}

// Eval64 returns a float64 Perlin noise value for the given coordinates.
func (n *Noise) Eval64(dim ...float64) float64 {
	switch len(dim) {
	case 1:
		return n.eval1D(dim[0])
	case 2:
		return n.eval2D(dim[0], dim[1])
	case 3:
		return n.eval3D(dim[0], dim[1], dim[2])
	}

	return 0
}

// SetEval32 is an implementation of noise.Generator interface. It will always
// return an error.
func (n *Noise) SetEval32(f func(seed int64, dim ...float32) float32) error {
	return errors.New("float32 evaluation function is already set")
}

// SetEval64 is an implementation of noise.Generator interface. It will always
// return an error.
func (n *Noise) SetEval64(f func(seed int64, dim ...float64) float64) error {
	return errors.New("float64 evaluation function is already set")
}

// ----------------------------------------------------------------------------
//  Methods (Private)
// ----------------------------------------------------------------------------

// eval1D generates float64 Perlin noise value from 1-dimensional coordinate.
func (n *Noise) eval1D(x float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise1D(x)
}

// eval2D generates float64 Perlin noise value from 2-dimensional coordinates.
func (n *Noise) eval2D(x, y float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise2D(x, y)
}

// eval3D generates float64 Perlin noise value from 3-dimensional coordinates.
func (n *Noise) eval3D(x, y, z float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise3D(x, y, z)
}
