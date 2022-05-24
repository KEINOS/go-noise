package perlin_test

import (
	"fmt"

	"github.com/KEINOS/go-noise/pkg/perlin"
)

func ExampleNew_eval32_one_dimmention() {
	const seed = 100

	p := perlin.New(seed)

	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Eval32(float32(x/100)))
	}

	// Output:
	// 0;0.0000
	// 1;-0.0026
	// 2;-0.0046
}

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

func ExampleNew_eval64_two_dimmention() {
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

func ExampleNew_eval64_three_dimmention() {
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
