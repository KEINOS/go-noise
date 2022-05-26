package custom_test

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/KEINOS/go-noise/pkg/custom"
)

// Example to use the custom noise generator for float32.
//
// This example uses the `custom.New` function to instantiate a new noise
// generator. Then the `custom.SetEval32` function is used to set the user-
// defined function.
//
//nolint:dupl // This example is intentionally duplicated.
func ExampleGenerator_SetEval32() {
	const seed = int64(12345)
	var rnd *rand.Rand

	// Instantiate a new noise generator.
	genNoise := custom.New(seed)

	// User custom function.
	// It generates pseudo-random numbers between -1 and 1.
	myCustomFunc := func(seed int64, dim ...float32) float32 {
		if rnd == nil {
			rnd = rand.New(rand.NewSource(seed))
		}

		for i := 0; i < len(dim); i++ {
			max := int(dim[i])

			for ii := 0; ii < max; ii++ {
				_ = rnd.Float32()
			}
		}

		// Generate a pseudo-random number.
		//nolint:gosec // Use of weak random number generation is intended for simple examples.
		v := rnd.Float32()

		return v*2 - 1 // Convert [0.0,1.0] to [-1.0,1.0]
	}

	// Set the user-defined function.
	// The function must return a value between -1 and 1. Same seed values should
	// generate the same pseudo-random sequence.
	if err := genNoise.SetEval32(myCustomFunc); err != nil {
		log.Fatal(err)
	}

	out := genNoise.Eval32(1.0, 2.0, 3.0)

	fmt.Printf("%f\n", out)

	// Output: -0.969880
}

// Example to use the custom noise generator for float64.
//
// This example uses the `custom.New` function to instantiate a new noise
// generator. Then the `custom.SetEval642` function is used to set the user-
// defined function.
//
//nolint:dupl // This example is intentionally duplicated.
func ExampleGenerator_SetEval64() {
	const seed = int64(12345)
	var rnd *rand.Rand

	// Instantiate a new noise generator.
	genNoise := custom.New(seed)

	// User custom function.
	myCustomFunc := func(seed int64, dim ...float64) float64 {
		if rnd == nil {
			rnd = rand.New(rand.NewSource(seed))
		}

		for i := 0; i < len(dim); i++ {
			max := int(dim[i])

			for ii := 0; ii < max; ii++ {
				_ = rnd.Float64()
			}
		}

		// Generate a pseudo-random number.
		//nolint:gosec // Use of weak random number generation is intended for simple examples.
		v := rnd.Float64()

		return v*2 - 1 // Convert [0.0,1.0] to [-1.0,1.0]
	}

	// Set the user-defined function.
	// The function must return a value between -1 and 1. Same seed values should
	// generate the same pseudo-random sequence.
	if err := genNoise.SetEval64(myCustomFunc); err != nil {
		log.Fatal(err)
	}

	out := genNoise.Eval64(1.0, 2.0, 3.0)

	fmt.Printf("%f\n", out)

	// Output: -0.969880
}
