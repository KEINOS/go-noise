package main

import (
	"log"

	"github.com/KEINOS/go-noise"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	// Image size of X and Y.
	width  = 500
	height = 500
)

// These values should be adjusted to obtain the appropriate dispersion for
// the image size.
const (
	// Seed is like pattern ID. If the seed values are the same, the noise
	// pattern will also be the same.
	seed = 100
	// The higher the smoother or less difference between the plotting points.
	smoothness = 100
)

func main() {
	{
		// Create Perlin noise generator with seed 100.
		gen, err := noise.New(noise.Perlin, seed)
		if err != nil {
			log.Fatal(err)
		}

		pathFile := "./2d_perlin.png"

		if err := GenGraphImage(gen, "Perlin Noise", pathFile); err != nil {
			log.Fatal(err)
		}
	}

	{
		// Create OpenSimplex noise generator with seed 100.
		gen, err := noise.New(noise.OpenSimplex, seed)
		if err != nil {
			log.Fatal(err)
		}

		pathFile := "./2d_opensimplex.png"

		if err := GenGraphImage(gen, "OpenSimplex Noise", pathFile); err != nil {
			log.Fatal(err)
		}
	}
}

// GenGraphImage generates a graph image with the specified generator and saves
// it to the specified path.
func GenGraphImage(gen noise.Noise, title string, pathFile string) error {
	// Create plotter.
	pltr := GenPlotter(title)

	// Create plotting function
	pltf := plotter.NewFunction(func(x float64) float64 {
		xx := x / smoothness

		// Generate noise at yy from xx.
		yy := gen.Eval64(xx)

		return Normalize(yy) // Return normalized value.
	})

	// Add the function to the plotter.
	pltr.Add(pltf)

	// Set loop range.
	pltr.X.Min = 0
	pltr.X.Max = float64(width)
	pltr.Y.Min = 0
	pltr.Y.Max = float64(height)

	// Save image
	if err := pltr.Save(3*vg.Inch, 3*vg.Inch, pathFile); err != nil {
		return err
	}

	return nil
}

// GenPlotter returns a plotter with the specified title and basic settings.
func GenPlotter(titile string) *plot.Plot {
	p := plot.New()

	p.Title.Text = titile
	p.X.Label.Text = "X axis"
	p.Y.Label.Text = "Y axis"

	return p
}

// Normalize converts the value between -1 and 1 to 0 and height.
func Normalize(v float64) float64 {
	return (v + 1) / 2 * float64(height)
}
