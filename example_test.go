package noise_test

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/KEINOS/go-noise"
)

func ExampleNew() {
	//nolint:gosec // Example value
	seed := rand.Int63()

	noise, err := noise.New(noise.OpenSimplex, seed)
	if err != nil {
		log.Fatal(err)
	}

	w, h := 100, 100

	heightmap := make([]float64, w*h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x / w)
			yFloat := float64(y / h)

			heightmap[(y*w)+x] = noise.Eval64(xFloat, yFloat)
		}
	}
}

// ----------------------------------------------------------------------------
//  Custom Noise (User-Defind Fucntion)
// ----------------------------------------------------------------------------

func ExampleNew_assign_user_defined_function() {
	const seed = 100

	// Create a noise generator.
	gen, err := noise.New(noise.Custom, seed)
	if err != nil {
		log.Fatal(err)
	}

	// rnd is a random number generator.
	var rnd *rand.Rand

	// Define user function to generate custom noise.
	// Here we define a function that returns a pseudo-random value from the
	// given seed and itereate 'dim' times.
	myFunc := func(seed int64, dim ...float32) float32 {
		if rnd == nil {
			//nolint:gosec // Use of weak random number generation is intended here.
			rnd = rand.New(rand.NewSource(seed))
		}

		for i := 0; i < len(dim); i++ {
			max := int(dim[i])

			for ii := 0; ii < max; ii++ {
				_ = rnd.Float32()
			}
		}

		// Generate a pseudo-random number.
		v := rnd.Float32()

		return v*2 - 1 // Convert [0.0,1.0] to [-1.0,1.0]
	}

	// Assign user-defined function
	if err := gen.SetEval32(myFunc); err != nil {
		log.Fatal(err)
	}

	// Get noise value at (1, 2, 3)
	fmt.Println(gen.Eval32(1, 2, 3))

	// Output: -0.1159181
}

// ----------------------------------------------------------------------------
//  OpenSimplex Noise
// ----------------------------------------------------------------------------

func ExampleNew_one_dimension_opensimplex_noise_in_float32() {
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

func ExampleNew_two_dimension_opensimplex_noise_in_float32() {
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

func ExampleNew_three_dimension_opensimplex_noise_in_float32() {
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

func ExampleNew_one_dimension_opensimplex_in_float64() {
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
	// 0.0000;0.0000
	// 0.1000;0.0950
	// 0.2000;0.1385
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
