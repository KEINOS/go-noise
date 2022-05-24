package opensimplex_test

import (
	"fmt"

	"github.com/KEINOS/go-noise/pkg/opensimplex"
)

// -----------------------------------------------------------------------------
//  Noise.Eval32 implementation test
// -----------------------------------------------------------------------------

func ExampleNew_eval32_one_dimmention() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for x := float32(0); x < 3; x++ {
		fmt.Printf("%0.4f\n", p.Eval32(x/smoothness))
	}

	// Output:
	// 0.0000
	// 0.0102
	// 0.0204
}

func ExampleNew_eval32_two_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for y := float32(0); y < 2; y++ {
		for x := float32(0); x < 2; x++ {
			fmt.Printf(
				"%0.4f\n",
				p.Eval32(x/smoothness, y/smoothness),
			)
		}
	}

	// Output:
	// 0000
	// 0.0170
	// -0.0068
	// 0.010
}

func ExampleNew_eval32_three_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for z := float32(0); z < 2; z++ {
		for y := float32(0); y < 2; y++ {
			for x := float32(0); x < 2; x++ {
				fmt.Printf(
					"%0.4f\n",
					p.Eval32(x/smoothness, y/smoothness, z/smoothness),
				)
			}
		}
	}

	// Output:
	// 0.0000
	// -0.0062
	// 0.0062
	// 0.0000
	// -0.0171
	// -0.0233
	// -0.0109
	// -0.0171
}

func ExampleNew_eval32_more_than_three_dimmentions() {
	const (
		seed = 100
	)

	p := opensimplex.New(seed)

	// It only supports up to three dimmentions.
	// OpenSimplex itself supports more but currently we strictly limit it to three.
	fmt.Printf("%0.0f", p.Eval32(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}

// -----------------------------------------------------------------------------
//  Noise.Eval64 implementation test
// -----------------------------------------------------------------------------

func ExampleNew_eval64_one_dimmention() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for x := float64(0); x < 3; x++ {
		fmt.Printf(
			"%0.4f\n",
			p.Eval64(x/smoothness),
		)
	}

	// Output:
	// 0.0000
	// 0.0102
	// 0.0204
}

func ExampleNew_eval64_two_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf(
					"%0.4f\n",
					p.Eval64(x/smoothness, y/smoothness, z/smoothness),
				)
			}
		}
	}

	// Output:
	// 0.0000
	// -0.0171
	// 0.0062
	// -0.0109
	// -0.0062
	// -0.0233
	// 0.0000
	// -0.0171
}

func ExampleNew_eval64_three_dimmentions() {
	const (
		seed       = 100
		smoothness = 100
	)

	p := opensimplex.New(seed)

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf(
					"%0.4f\n",
					p.Eval64(x/smoothness, y/smoothness, z/smoothness),
				)
			}
		}
	}

	// Output:
	// 0.0000
	// -0.0171
	// 0.0062
	// -0.0109
	// -0.0062
	// -0.0233
	// 0.0000
	// -0.0171
}

func ExampleNew_eval64_more_than_three_dimmentions() {
	const (
		seed = 100
	)

	p := opensimplex.New(seed)

	// It only supports up to three dimmentions.
	// OpenSimplex itself supports more but currently we strictly limit it to three.
	fmt.Printf("%0.0f", p.Eval64(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}
