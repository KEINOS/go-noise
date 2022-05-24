/*
Package perlin is a wrapper of github.com/aquilax/go-perlin.

Which implements github.com/KEINOS/go-noise/noise interface.
*/
package perlin

import goperlin "github.com/aquilax/go-perlin"

const (
	Alpha     = 2.       // Alpha is the default smoothness of Perlin noise.
	Beta      = 2.       // Beta is the default scale of Perlin noise.
	Iteration = int32(3) // Iteration is the default number of iterations of Perlin noise.
)

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New returns a seeded Perlin noise instance.
func New(seed int64) *Noise {
	return &Noise{
		Smoothness: Alpha,
		Scale:      Beta,
		Seed:       int64(seed),
		Iteration:  int32(Iteration),
	}
}

// ----------------------------------------------------------------------------
//  Type: Noise
// ----------------------------------------------------------------------------

type Noise struct {
	// Smoothness is the weight when the sum is formed. Which is the alpha value
	// in Ken Perlin's article. Default is 2. The smaller the number, the more
	//noise is generated.
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
//  Methods
// ----------------------------------------------------------------------------

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

func (n *Noise) eval1D(x float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise1D(x)
}

// Eval2D generates 2-dimensional Perlin Noise value.
func (n *Noise) eval2D(x, y float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise2D(x, y)
}

// Eval3D generates 3-dimensional Perlin Noise value.
func (n *Noise) eval3D(x, y, z float64) float64 {
	p := goperlin.NewPerlin(n.Smoothness, n.Scale, n.Iteration, n.Seed)

	return p.Noise3D(x, y, z)
}
