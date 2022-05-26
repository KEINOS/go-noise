package opensimplex_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-noise/pkg/opensimplex"
)

// -----------------------------------------------------------------------------
//  Generator.Eval32 implementation test
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
	// 0.0000
	// 0.0170
	// -0.0068
	// 0.0102
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
	const seed = 100

	p := opensimplex.New(seed)

	// It only supports up to three dimmentions.
	// OpenSimplex itself supports more but currently we strictly limit it to three.
	fmt.Printf("%0.0f", p.Eval32(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}

// -----------------------------------------------------------------------------
//  Generator.Eval64 implementation test
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
	const seed = 100

	p := opensimplex.New(seed)

	// It only supports up to three dimmentions.
	// OpenSimplex itself supports more but currently we strictly limit it to three.
	fmt.Printf("%0.0f", p.Eval64(0.0001, 0.0001, 0.0001, 0.0001))

	// Output: 0
}

func ExampleGenerator_SetEval32() {
	const seed = 100

	p := opensimplex.New(seed)

	// User-defined functions cannot be assigned to OpenSimplex types. Use Custom type instead.
	err := p.SetEval32(func(seed int64, dim ...float32) float32 {
		return 0
	})

	if err == nil {
		log.Fatal("OpenSimplex type should return an error on SetEval32")
	}

	fmt.Println(err.Error())

	// Output:
	// float32 evaluation function is already set. You can not set custom function in OpenSimplex type
}

func ExampleGenerator_SetEval64() {
	const seed = 100

	p := opensimplex.New(seed)

	// User-defined functions cannot be assigned to OpenSimplex types. Use Custom type instead.
	err := p.SetEval64(func(seed int64, dim ...float64) float64 {
		return 0
	})

	if err == nil {
		log.Fatal("OpenSimplex type should return an error on SetEval32")
	}

	fmt.Println(err.Error())

	// Output:
	// float64 evaluation function is already set. You can not set custom function in OpenSimplex type
}
