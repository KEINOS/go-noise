package perlin_test

import (
	"fmt"

	"github.com/KEINOS/go-noise/pkg/perlin"
)

// ----------------------------------------------------------------------------
//  Noise.Eval32 implementation test
// ----------------------------------------------------------------------------

func ExampleNew_eval32_one_dimmention() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := perlin.New(seed)

	for x := float32(0); x < 3; x++ {
		fmt.Printf(
			"%0.0f;%0.4f\n",
			x,
			p.Eval32(x/smoothness),
		)
	}

	// Output:
	// 0;0.0000
	// 1;-0.0026
	// 2;-0.0046
}

func ExampleNew_eval32_two_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := perlin.New(seed)

	for y := float32(0); y < 3; y++ {
		for x := float32(0); x < 3; x++ {
			fmt.Printf(
				"%0.0f;%0.0f; %0.4f\n",
				x, y,
				p.Eval32(x/smoothness, y/smoothness),
			)
		}
	}

	// Output:
	// 0;0; 0.0000
	// 1;0; -0.0285
	// 2;0; -0.0602
	// 0;1; -0.0152
	// 1;1; -0.0437
	// 2;1; -0.0753
	// 0;2; -0.0328
	// 1;2; -0.0612
	// 2;2; -0.0927
}

func ExampleNew_eval32_three_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := perlin.New(seed)

	for z := float32(0); z < 2; z++ {
		for y := float32(0); y < 2; y++ {
			for x := float32(0); x < 2; x++ {
				fmt.Printf(
					"%0.0f;%0.0f;%0.0f; %0.4f\n",
					x, y, z,
					p.Eval32(x/smoothness, y/smoothness, z/smoothness),
				)
			}
		}
	}

	// Output:
	// 0;0;0; 0.0000
	// 1;0;0; -0.0146
	// 0;1;0; -0.0104
	// 1;1;0; -0.0249
	// 0;0;1; 0.0257
	// 1;0;1; 0.0112
	// 0;1;1; 0.0154
	// 1;1;1; 0.0009
}

func ExampleNew_eval32_more_than_three_dimmentions() {
	const (
		seed = 100
	)

	p := perlin.New(seed)

	// It only supports up to three dimmentions.
	fmt.Printf("%0.0f", p.Eval32(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}

// ----------------------------------------------------------------------------
//  Noise.Eval64 implementation test
// ----------------------------------------------------------------------------

func ExampleNew_eval64_one_dimmention() {
	const seed = 100

	p := perlin.New(seed)

	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Eval64(x/10))
	}

	// Output:
	// 0;0.0000
	// 1;-0.0086
	// 2;-0.0017
}

func ExampleNew_eval64_two_dimmentions() {
	const seed = 100

	p := perlin.New(seed)

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, p.Eval64(x/10, y/10))
		}
	}

	// Output:
	// 0;0;0.0000
	// 0;1;-0.2002
	// 1;0;-0.3389
	// 1;1;-0.5045
}

func ExampleNew_eval64_three_dimmentions() {
	const seed = 100

	p := perlin.New(seed)

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("%0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, p.Eval64(x/10, y/10, z/10))
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

func ExampleNew_eval64_more_than_three_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := perlin.New(seed)

	// It only supports up to three dimmentions.
	fmt.Printf("%0.0f", p.Eval64(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}
