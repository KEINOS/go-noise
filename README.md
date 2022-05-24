# Go-Noise

[go-noise](https://github.com/KEINOS/go-noise) is a Go package for generating a gradient noise between -1 and 1.

It's a wrapper of [github.com/aquilax/go-perlin](https://github.com/aquilax/go-perlin) and [github.com/ojrac/opensimplex-go](https://github.com/ojrac/opensimplex-go) to make it easy to use.

## 1D Example

In the example below, 1 dimmentional method `Noise.Eval64(x)` was used to generate the noise value `y` at the position `x`.

```go
// seed := 100       // noise pattern ID
// smoothness := 100 // noise smoothness
//
// noiseType := noise.Perlin
// noiseType := noise.OpenSimplex
n, err := noise.New(noiseType, seed)

yy := n.Eval64(x / smoothness) // yy is between -1.0 and 1.0 of float64
y := (yy + 1) / 2 * 500        // y is between 0 and 500
```

![](./_example/2d/2d_perlin.png)
![](./_example/2d/2d_opensimplex.png)

- [Source](./_example/2d)

## 2D Example

```go
// Obtain the noise value at the position (x, y)
n, err := noise.New(noiseType, seed)
v := n.Eval64(x, y) // v is between -1.0 and 1.0 of float64
```

To create a 2D image, plot the `v` value at the position `(x, y)` in the 2D space. The 2D image example is equivalent to a frame of the 3D image example below.

## 3D Example

In the example below, three dimmentional method `noise.Eval64(x, y, z)` was used to generate the noise value `x, y` at the position `z`.

The X and Y axes are the 2D image and the Z axis is the "time", or animation frame, of the 3D noise sample.

### Perlin Noise Sample

```go
// Obtain the noise value at the position (x, y, z)
n, err := noise.New(noise.Perlin, seed)
v := n.Eval64(x, y, z) // v is between -1.0 and 1.0 of float64
```

![](./_example/3d/animation_perlin.gif)

- [Source](./_example/3d)

### OpenSimplex Noise Sample

```go
// Obtain the noise value at the position (x, y, z)
n, err := noise.New(noise.OpenSimplex, seed)
v := n.Eval64(x, y, z) // v is between -1.0 and 1.0 of float64
```

![](./_example/3d/animation_opensimplex.gif)

- [Source](./_example/3d)

### Note

This package does NOT support more than 3 dimensions. If more than 3 dimentions were given, such as `Noise.Eval64(w, x, y, z)`, it will retrun a `0` (zero) value.

## Usage

### Require module

```bash
go get "github.com/KEINOS/go-noise"
```

### Constructor

```go
import "github.com/KEINOS/go-noise"

// Seed is like pattern ID.
// If the seed values are the same, the noise pattern will also be the
// same.
const seed = 100

// Noise generator for Perlin noise
genNoise, err := noise.New(noise.Perlin, seed)
```
```go
import "github.com/KEINOS/go-noise"

// Seed is like pattern ID.
// If the seed values are the same, the noise pattern will also be the
// same.
const seed = 100

// Noise generator for OpenSimplex noise
genNoise, err := noise.New(noise.OpenSimplex, seed)
```

### Methods

```go
a := genNoise.Eval32(x)       // 1D noise. Generate noise at x.
b := genNoise.Eval32(x, y)    // 2D noise. Generate noise at x, y.
c := genNoise.Eval32(x, y, z) // 3D noise. Generate noise at x, y, z.

// a, b, c, x, y, z are float32.
// Noises a, b, c are between -1.0 and 1.0.
```
```go
a := genNoise.Eval64(x)       // 1D noise. Generate noise at x.
b := genNoise.Eval64(x, y)    // 2D noise. Generate noise at x, y.
c := genNoise.Eval64(x, y, z) // 3D noise. Generate noise at x, y, z.

// a, b, c, x, y, z are float64.
// Noises a, b, c are between -1.0 and 1.0.
```

### Brief Example

```go
import "github.com/KEINOS/go-noise"

const seed = 100
const steps = 5
const zoom = 25
const frame = 50

// Create new noise generator of Perlin type
genNoise, err := noise.New(noise.Perlin, seed)

if err != nil {
    // error handle
}

for z := 0; z < frame; z++ {
    zz := float64(z) / steps

    /* Here create a new image of a frame */

    for y := 0; y < height; y++ {
        yy := float64(y) / zoom

        for x := 0; x < width; x++ {
            xx := float64(x) / zoom

            // n is a float64 value between -1 and 1
            n := genNoise.Eval64(xx, yy, zz)

            // Convert n to 0-255 scale
            grayColor := ((1. - in) / 2.) * 255.

            pixelToPlot := color.Gray{Y: uint8(grayColor)}

            /* Here plot the pixel to the current image */
        }
    }

    /* Here save the current frame/image to a file */
}

/* Here animate the frames if you want */
```

- [Complete Source](./_example/3d)

#### Advanced

To change `alpha` and `beta` values of Perling noise.

```go
alpha := 0.5
beta := 0.5

n, err := noise.New(noise.Perlin, seed)
n.Smoothness = alpha
n.Scale = beta
```
