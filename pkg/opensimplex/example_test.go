package opensimplex_test

import (
	"fmt"

	"github.com/KEINOS/go-noise/pkg/opensimplex"
)

func ExampleNew_one_dimmention() {
	const seed = 100

	p := opensimplex.New(seed)

	for x := 0.; x < 3; x++ {
		fmt.Printf("float32: %0.0f;%0.4f\n", x, p.Eval32(float32(x/10)))
	}

	for x := 0.; x < 3; x++ {
		fmt.Printf("float64: %0.0f;%0.4f\n", x, p.Eval64(x/10))
	}

	// Output:
	// float32: 0;0.0000
	// float32: 1;0.0950
	// float32: 2;0.1385
	// float64: 0;0.0000
	// float64: 1;0.0950
	// float64: 2;0.1385
}

func ExampleNew_three_dimmention() {
	const seed = 100

	p := opensimplex.New(seed)

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("float32: %0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, p.Eval32(float32(x/10), float32(y/10), float32(z/10)))
			}
		}
	}

	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("float64: %0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, p.Eval64(x/10, y/10, z/10))
			}
		}
	}

	// Output:
	// float32: 0;0;0;0.0000
	// float32: 0;0;1;-0.1672
	// float32: 0;1;0;0.0607
	// float32: 0;1;1;-0.1040
	// float32: 1;0;0;-0.0611
	// float32: 1;0;1;-0.2227
	// float32: 1;1;0;0.0003
	// float32: 1;1;1;-0.1585
	// float64: 0;0;0;0.0000
	// float64: 0;0;1;-0.1672
	// float64: 0;1;0;0.0607
	// float64: 0;1;1;-0.1040
	// float64: 1;0;0;-0.0611
	// float64: 1;0;1;-0.2227
	// float64: 1;1;0;0.0003
	// float64: 1;1;1;-0.1585
}
