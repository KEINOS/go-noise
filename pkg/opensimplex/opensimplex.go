package opensimplex

import orig "github.com/ojrac/opensimplex-go"

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

type Noise struct {
	// Seed holds the seed value for the noise.
	Seed int64
}

// ----------------------------------------------------------------------------
//  Methods
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

// eval1D32 generates 1-dimensional OpenSimplex Noise value.
func (n *Noise) eval1D32(x float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval2(x, x)
}

// eval2D32 generates 2-dimensional OpenSimplex Noise value.
func (n *Noise) eval2D32(x, y float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval2(x, y)
}

// eval3D32 generates 3-dimensional OpenSimplex Noise value.
func (n *Noise) eval3D32(x, y, z float32) float32 {
	p := orig.New32(n.Seed)

	return p.Eval3(x, y, z)
}

// eval1D64 generates 1-dimensional OpenSimplex Noise value.
func (n *Noise) eval1D64(x float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval2(x, x)
}

// eval2D64 generates 2-dimensional OpenSimplex Noise value.
func (n *Noise) eval2D64(x, y float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval2(x, y)
}

// eval3D64 generates 3-dimensional OpenSimplex Noise value.
func (n *Noise) eval3D64(x, y, z float64) float64 {
	p := orig.New(n.Seed)

	return p.Eval3(x, y, z)
}
