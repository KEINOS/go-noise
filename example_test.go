package noise_test

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/KEINOS/go-noise"
)

func ExampleNew() {
	noise, err := noise.New(noise.OpenSimplex, rand.Int63())
	if err != nil {
		log.Fatal(err)
	}

	w, h := 100, 100
	heightmap := make([]float64, w*h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)

			heightmap[(y*w)+x] = noise.Eval64(xFloat, yFloat)
		}
	}
}

// ----------------------------------------------------------------------------
//  OpenSimplex Noise
// ----------------------------------------------------------------------------

func ExampleNew_one_dimention_opensimplex_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.OpenSimplex, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 3; x++ {
		fmt.Printf(
			"%0.0f;%0.4f\n",
			x,
			gen.Eval32(x/10),
		)
	}

	// Output:
	// 0;0.0000
	// 1;0.0950
	// 2;0.1385
}

func ExampleNew_two_dimention_opensimplex_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.OpenSimplex, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 2; x++ {
		for y := float32(0); y < 2; y++ {
			fmt.Printf(
				"%0.0f;%0.0f;%0.4f\n",
				x, y,
				gen.Eval32(x/10, y/10),
			)
		}
	}

	// Output:
	// 0;0;0.0000
	// 0;1;-0.0673
	// 1;0;0.1664
	// 1;1;0.0950
}

func ExampleNew_three_dimention_opensimplex_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.OpenSimplex, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 2; x++ {
		for y := float32(0); y < 2; y++ {
			for z := float32(0); z < 2; z++ {
				fmt.Printf(
					"%0.0f;%0.0f;%0.0f;%0.4f\n",
					x, y, z,
					gen.Eval32(x/10, y/10, z/10),
				)
			}
		}
	}

	// Output:
	// 0;0;0;0.0000
	// 0;0;1;-0.1672
	// 0;1;0;0.0607
	// 0;1;1;-0.1040
	// 1;0;0;-0.0611
	// 1;0;1;-0.2227
	// 1;1;0;0.0003
	// 1;1;1;-0.1585
}

func ExampleNew_one_dimention_opensimplex_in_float64() {
	const seed = 100

	gen, err := noise.New(noise.OpenSimplex, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float64(0); x < 3; x++ {
		fmt.Printf("%0.4f;%0.4f\n",
			x/10,
			gen.Eval64(x/10),
		)
	}

	// Output:
	// 0;0.0000
	// 1;0.0950
	// 2;0.1385
}

// ----------------------------------------------------------------------------
//  Perlin Noise
// ----------------------------------------------------------------------------

func ExampleNew_one_dimension_perlin_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, gen.Eval32(x/100))
	}

	// Output:
	// 0;0.0000
	// 1;-0.0026
	// 2;-0.0046
}

func ExampleNew_two_dimension_perlin_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 2; x++ {
		for y := float32(0); y < 2; y++ {
			fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, gen.Eval32(x/10, y/10))
		}
	}

	// Output:
	// 0;0;0.0000
	// 0;1;-0.2002
	// 1;0;-0.3389
	// 1;1;-0.5045
}

func ExampleNew_three_dimension_perlin_noise_in_float32() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float32(0); x < 2; x++ {
		for y := float32(0); y < 2; y++ {
			for z := float32(0); z < 2; z++ {
				fmt.Printf("%0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, gen.Eval32(x/10, y/10, z/10))
			}
		}
	}

	// Output:
	// 0;0;0;0.0000
	// 0;0;1;0.2616
	// 0;1;0;-0.0755
	// 0;1;1;0.2020
	// 1;0;0;-0.2138
	// 1;0;1;0.0616
	// 1;1;0;-0.2208
	// 1;1;1;0.0304
}

func ExampleNew_one_dimension_perlin_noise_in_float64() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float64(0); x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, gen.Eval64(x/10))
	}

	// Output:
	// 0;0.0000
	// 1;-0.0086
	// 2;-0.0017
}

func ExampleNew_two_dimension_perlin_noise_in_float64() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float64(0); x < 2; x++ {
		for y := float64(0); y < 2; y++ {
			fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, gen.Eval64(x/10, y/10))
		}
	}

	// Output:
	// 0;0;0.0000
	// 0;1;-0.2002
	// 1;0;-0.3389
	// 1;1;-0.5045
}

func ExampleNew_three_dimension_perlin_noise_in_float64() {
	const seed = 100

	gen, err := noise.New(noise.Perlin, seed)
	if err != nil {
		log.Fatal(err)
	}

	for x := float64(0); x < 2; x++ {
		for y := float64(0); y < 2; y++ {
			for z := float64(0); z < 2; z++ {
				fmt.Printf("%0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, gen.Eval64(x/10, y/10, z/10))
			}
		}
	}

	// Output:
	// 0;0;0;0.0000
	// 0;0;1;0.2616
	// 0;1;0;-0.0755
	// 0;1;1;0.2020
	// 1;0;0;-0.2138
	// 1;0;1;0.0616
	// 1;1;0;-0.2208
	// 1;1;1;0.0304
}
