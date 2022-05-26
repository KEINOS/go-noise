package opensimplex

import (
	"github.com/pkg/errors"

	orig "github.com/ojrac/opensimplex-go"
)

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a seeded Perlin noise instance.
func New(seed int64) *Noise {
	return &Noise{
		Seed: seed,
	}
}

// ----------------------------------------------------------------------------
//  Type: Noise
// ----------------------------------------------------------------------------

// Noise holds parameter values for OpenSimplex noise. It is an implementation
// of Noise interface.
type Noise struct {
	// Seed holds the seed value for the noise.
	Seed int64
}

// ----------------------------------------------------------------------------
//  Methods (Public)
// ----------------------------------------------------------------------------

// Eval32 returns a float32 noise value for the given coordinates.
func (n *Noise) Eval32(dim ...float32) float32 {
	switch len(dim) {
	case 1:
		return n.eval1D32(dim[0])
	case 2:
		return n.eval2D32(dim[0], dim[1])
	case 3:
		return n.eval3D32(dim[0], dim[1], dim[2])
	}

	return 0
}

// Eval64 returns a float64 noise value for the given coordinates.
func (n *Noise) Eval64(dim ...float64) float64 {
	switch len(dim) {
	case 1:
		return n.eval1D64(dim[0])
	case 2:
		return n.eval2D64(dim[0], dim[1])
	case 3:
		return n.eval3D64(dim[0], dim[1], dim[2])
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

// eval1D32 generates float32 OpenSimplex noise value from 1-dimensional coordinate.
func (n *Noise) eval1D32(x float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval2(x, x)
}

// eval2D32 generates float32 OpenSimplex noise value from 2-dimensional coordinates.
func (n *Noise) eval2D32(x, y float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval2(x, y)
}

// eval3D32 generates float32 OpenSimplex noise value from 3-dimensional coordinates.
func (n *Noise) eval3D32(x, y, z float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval3(x, y, z)
}

// eval1D64 generates float64 OpenSimplex noise value from 1-dimensional coordinate.
func (n *Noise) eval1D64(x float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval2(x, x)
}

// eval2D64 generates float64 OpenSimplex noise value from 2-dimensional coordinates.
func (n *Noise) eval2D64(x, y float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval2(x, y)
}

// eval3D64 generates float64 OpenSimplex noise value from 3-dimensional coordinates.
func (n *Noise) eval3D64(x, y, z float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval3(x, y, z)
}
