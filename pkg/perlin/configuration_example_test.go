package perlin_test

import (
	"fmt"

	"github.com/KEINOS/go-noise/pkg/perlin"
)

func ExampleNew_configuration() {
	const seed = 100

	generator := perlin.New(seed)
	generator.Smoothness = 1.5
	generator.Scale = 1.5

	fmt.Printf("%.4f\n", generator.Eval64(0.1, 0.2))

	// Output: -0.7206
}
