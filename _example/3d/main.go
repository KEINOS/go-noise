package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"

	"github.com/KEINOS/go-noise"
	"github.com/pkg/errors"
)

const (
	// Image size of X and Y.
	width  = 250
	height = 250
	// Z-axis = number of frames of the animation.
	depth = 50
)

// These values should be adjusted to obtain the appropriate dispersion for
// the image size.
const (
	// Create noise generator with seed 100. Same seed will generate the same
	// noise pattern.
	seed = 100

	// The bigger the zoom and less randomness.
	zoom = 25
	// Smoothness between frames. Smaller value is smoother but little difference
	// between frames.
	steps = 5
)

// ----------------------------------------------------------------------------
//  Main
// ----------------------------------------------------------------------------

func main() {
	{
		fmt.Println("Creating Perlin noise image ...")

		gen, err := noise.New(noise.Perlin, seed)
		if err != nil {
			log.Fatal(err)
		}

		if err := GenNoiseImage(gen, "animation_perlin.gif"); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nPerlin Noise image created.\n")
	}

	{
		fmt.Println("Creating OpenSimplex noise image ...")

		gen, err := noise.New(noise.OpenSimplex, seed)
		if err != nil {
			log.Fatal(err)
		}

		if err := GenNoiseImage(gen, "animation_opensimplex.gif"); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nOpenSimplex Noise image created.\n")
	}

	fmt.Printf("\nDone!\n")

}

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// CreatePaletteGray creates a color.Palette of 256 gray-scale colors.
func CreatePaletteGray() color.Palette {
	grayPalet := make(color.Palette, 256)

	for c := 0; c < 256; c++ {
		grayPalet[c] = color.Gray{Y: uint8(c)}
	}

	return grayPalet
}

func GenNoiseImage(gen noise.Noise, pathFile string) error {
	// Define boundary of image
	myBound := image.Rectangle{Max: image.Point{X: width, Y: height}}
	// Create a color palette
	grayPalet := CreatePaletteGray()
	// Create a new image
	outGif := &gif.GIF{}

	// Z-axis is the frame number
	for z := 0; z < depth; z++ {
		zz := float64(z) / steps
		canvas := image.NewPaletted(myBound, grayPalet)

		// Create frame image
		for y := 0; y < height; y++ {
			yy := float64(y) / zoom

			for x := 0; x < width; x++ {
				xx := float64(x) / zoom

				// Generate noise at xx, yy of layer zz
				grayC := NoiseToUint8(
					gen.Eval64(xx, yy, zz),
				)

				canvas.SetColorIndex(x, y, grayC) // Plot the pixel
			}
		}

		fmt.Printf("Frame: %d\r", z+1)
		outGif.Image = append(outGif.Image, canvas)
		outGif.Delay = append(outGif.Delay, 0)
	}

	// Export image
	if err := SaveImgGIF(outGif, pathFile); err != nil {
		return errors.Wrap(err, "failed to save image")
	}

	return nil
}

// NoiseToUint8 converts the "in" value to uint8.
//
// The "in" is the noise value which is between -1 and 1 of float64. This function
// converts it to a range of 0 to 255.
func NoiseToUint8(in float64) uint8 {
	out := ((1. - in) / 2.) * 255. // in: [-1, 1] --> out: [0, 255]

	return uint8(out)
}

func SaveImgGIF(img *gif.GIF, pathFile string) error {
	f, _ := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	if err := gif.EncodeAll(f, img); err != nil {
		return errors.Wrap(err, "failed to save image")
	}

	return nil
}
