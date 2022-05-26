package noise

import (
	"github.com/KEINOS/go-noise/pkg/custom"
	"github.com/KEINOS/go-noise/pkg/opensimplex"
	"github.com/KEINOS/go-noise/pkg/perlin"
	"github.com/pkg/errors"
)

// Algo represents the algorithm of the noise to generate.
type Algo int

const (
	// Unknown noise type.
	Unknown Algo = iota
	// Perlin noise type.
	Perlin
	// OpenSimplex noise type.
	OpenSimplex
	// Custom uses the user-defined function to generate noise.
	Custom
)

// ----------------------------------------------------------------------------
//  Types
// ----------------------------------------------------------------------------

// Generator is an interface for noise generator.
type Generator interface {
	// Eval32 returns a float32 noise value at given coordinates. The maximum
	// number of arguments is three.
	//
	// Example:
	//   Eval32(x)
	//   Eval32(x, y)
	//   Eval32(x, y, z)
	//
	// Implementations of this method must return the same value if the seed
	// value is the same.
	Eval32(dim ...float32) float32
	// Eval64 returns a float64 noise value at given coordinates. The maximum
	// number of arguments is three.
	//
	// Example:
	//   Eval64(x)
	//   Eval64(x, y)
	//   Eval64(x, y, z)
	//
	// Implementations of this method must return the same value if the seed
	// value is the same.
	Eval64(dim ...float64) float64
	SetEval32(f func(seed int64, dim ...float32) float32) error
	SetEval64(f func(seed int64, dim ...float64) float64) error
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a new noise generator.
func New(noiseType Algo, seed int64) (Generator, error) {
	switch noiseType {
	case Perlin:
		return perlin.New(seed), nil
	case OpenSimplex:
		return opensimplex.New(seed), nil
	case Custom:
		return custom.New(seed), nil
	}

	return nil, errors.New("unknown noise type")
}
