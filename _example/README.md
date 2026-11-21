# Basic Usage Example Sources

## 2D

This example plots one-dimensional noise as a 2D line graph. It uses Go 1.22 or later. The example module declares Gonum Plot and uses the go-noise source from this checkout. The custom noise plot gives the same value each time you evaluate the same seed and coordinates.

Run from the repository root:

```sh
(cd _example/2d && go run -mod=readonly .)
```

The command writes `2d_perlin.png`, `2d_opensimplex.png`, and `2d_pseudorandom.png` in `_example/2d`. It overwrites any existing files with these names.

## 3D

Example of 2 dimmension image with time lapse as the 3rd dimmension.

- Run: `cd ./_example/3d && go run .`

## Advanced Configuration

Perlin noise uses a default smoothness (`alpha`) of `2.0`, scale (`beta`) of `2.0`, and iteration count of `3`.

Use `perlin.New` to access the Perlin-specific `Smoothness` and `Scale` fields on `*perlin.Generator`.

```go
package main

import (
    "fmt"

    "github.com/KEINOS/go-noise/pkg/perlin"
)

func main() {
    const seed = 100

    generator := perlin.New(seed)
    generator.Smoothness = 1.5
    generator.Scale = 1.5

    fmt.Printf("%.4f\n", generator.Eval64(0.1, 0.2))
}
```

The example prints `-0.7206`. The [runnable Go example](../pkg/perlin/configuration_example_test.go) checks this configuration with `go test ./pkg/perlin` from the repository root.
