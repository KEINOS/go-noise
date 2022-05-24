package noise

import (
	"github.com/KEINOS/go-noise/pkg/opensimplex"
	"github.com/KEINOS/go-noise/pkg/perlin"
	"github.com/pkg/errors"
)

// Algo represents a noise type of the algorithm to generate noise.
type Algo int

const (
	// Unknown noise type.
	Unknown Algo = iota
	// Perlin noise type.
	Perlin
	// OpenSimplex noise type.
	OpenSimplex
)

// ----------------------------------------------------------------------------
//  Types
// ----------------------------------------------------------------------------

// Noise is an interface for noise generator.
type Noise interface {
	// Eval32 returns a float32 noise value at given coordinates. The maximum
	// number of arguments is three.
	//
	// Example:
	//   Eval32(x)
	//   Eval32(x, y)
	//   Eval32(x, y, z)
	Eval32(dim ...float32) float32
	// Eval64 returns a float64 noise value at given coordinates. The maximum
	// number of arguments is three.
	//
	// Example:
	//   Eval64(x)
	//   Eval64(x, y)
	//   Eval64(x, y, z)
	Eval64(dim ...float64) float64
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a new noise generator.
func New(noiseType Algo, seed int64) (Noise, error) {
	switch noiseType {
	case Perlin:
		return perlin.New(seed), nil
	case OpenSimplex:
		return opensimplex.New(seed), nil
	}

	return nil, errors.New("unknown noise type")
}
